#!/usr/bin/env python3
"""Append a job application record to ~/Desktop/jobs.csv."""

import csv
import os
import sys
from datetime import datetime

CSV_PATH = os.path.expanduser("~/Desktop/jobs.csv")
HEADERS = ["Name", "Link", "Applied?", "Date", "Comments"]


def main() -> None:
    if len(sys.argv) < 4:
        print(f"Usage: {sys.argv[0]} <name> <link> <applied?> [comments]", file=sys.stderr)
        sys.exit(1)

    name, link, applied = sys.argv[1:4]
    comments = sys.argv[4] if len(sys.argv) > 4 else ""
    date = datetime.now().strftime("%Y-%m-%d")

    file_exists = os.path.isfile(CSV_PATH)

    with open(CSV_PATH, "a", newline="", encoding="utf-8") as f:
        writer = csv.writer(f)
        if not file_exists:
            writer.writerow(HEADERS)
        writer.writerow([name, link, applied, date, comments])

    print(f"Recorded job '{name}' -> {CSV_PATH}")


if __name__ == "__main__":
    main()
