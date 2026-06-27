#!/usr/bin/env python3
import json
import re
import sys
import urllib.error
import urllib.request
from datetime import datetime

VALID_INDUSTRIES = {"Tech", "Health Care", "Retail", "Finance", "Gig", "Other"}
VALID_STATUSES = {
    "Not Started",
    "Applied Only",
    "Applied + Emailed",
    "Applied + Called",
    "Applied + Emailed + Called",
    "Interview!",
    "Done",
}
SCRIPT_URL = "https://script.google.com/macros/s/AKfycbwRQ52XCi5htaaHLO1Laizu8-pyYFKI0GEWELSnJHsP1CBDc-9OxNlkWGhlG-8l8tDxIQ/exec"


# Validation
def validate_url(url):
    if not url or not re.match(r"^https?://", url):
        raise ValueError("Link must be a valid URL starting with http:// or https://")
    return url


def validate_optional_email(email):
    """
    "@" and "." must exist somewhere in the email
    """
    if not email:
        return None
    if "@" not in email or "." not in email:
        raise ValueError("Email must be a valid email address")
    return email


def validate_optional_phone(phone):
    if not phone:
        return None

    digits = re.sub(r"[^\d]", "", phone)
    if not (10 <= len(digits) <= 15):
        raise ValueError("Phone number must be 10-15 digits")

    return digits


def validate_industry(industry):
    if industry not in VALID_INDUSTRIES:
        raise ValueError(
            f"Industry must be one of: {', '.join(sorted(VALID_INDUSTRIES))}"
        )
    return industry


def validate_status(status):
    # Case-insensitive matching
    status_lower = status.lower()
    for valid_status in VALID_STATUSES:
        if valid_status.lower() == status_lower:
            return valid_status
    raise ValueError(f"Status must be one of: {', '.join(sorted(VALID_STATUSES))}")


# CRUD Method
class NoRedirectHandler(urllib.request.HTTPRedirectHandler):
    def redirect_request(self, req, fp, code, msg, headers, newurl):
        return None


def post_to_sheet(data):
    body = json.dumps(data).encode("utf-8")
    req = urllib.request.Request(
        SCRIPT_URL,
        data=body,
        headers={"Content-Type": "application/json"},
        method="POST",
    )

    opener = urllib.request.build_opener(NoRedirectHandler)

    try:
        opener.open(req)
        print("Error: expected redirect, got none", file=sys.stderr)
        return 1
    except urllib.error.HTTPError as e:
        if e.code != 302:
            print(f"Error: HTTP {e.code} - {e.reason}", file=sys.stderr)
            return 1
        location = e.headers.get("Location")
        if not location:
            print("Error: no redirect Location header", file=sys.stderr)
            return 1

        # GET the redirect URL to get the actual response
        with urllib.request.urlopen(location) as response:
            result = response.read().decode("utf-8")
            if '"error"' in result:
                print(f"Fail: {result}", file=sys.stderr)
                return 1
            print(f"Success: {result}")
            return 0
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        return 1


# Entry
def main():
    if len(sys.argv) < 5:
        print(
            "Usage: track-job <companyName> <link> <industry> <status> [email] [phone] [notes]",
            file=sys.stderr,
        )
        print("\nRequired:", file=sys.stderr)
        print("\tcompanyName: Company name", file=sys.stderr)
        print("\tlink: Job posting URL", file=sys.stderr)
        print(
            "\tindustry: Tech, Health Care, Retail, Finance, Gig, Other",
            file=sys.stderr,
        )
        print(
            "\tstatus: Not Started, Applied Only, Applied + Emailed, Applied + Called, Applied + Emailed + Called, Interview!, Done",
            file=sys.stderr,
        )
        print("\nOptional:", file=sys.stderr)
        print("\temail: Employer contact email", file=sys.stderr)
        print("\tphone: Contact phone number", file=sys.stderr)
        print("\tnotes: Free-form notes", file=sys.stderr)
        return 1

    company_name = sys.argv[1]
    link = validate_url(sys.argv[2])
    industry = sys.argv[3]
    status = sys.argv[4]
    email = sys.argv[5] if len(sys.argv) > 5 else None
    phone = sys.argv[6] if len(sys.argv) > 6 else None
    notes = "".join(sys.argv[7:]) if len(sys.argv) > 7 else None

    if email is not None:
        email = validate_optional_email(email)

    if phone is not None:
        phone = validate_optional_phone(phone)

    industry = validate_industry(industry)
    status = validate_status(status)

    today = datetime.now().strftime("%Y-%m-%d")

    data = {
        "action": "create",
        "companyName": company_name,
        "link": link,
        "dateApplied": today,
        "industry": industry,
        "phoneNumber": phone,
        "email": email,
        "status": status,
        "notes": notes,
    }

    return post_to_sheet(data)


if __name__ == "__main__":
    sys.exit(main())
