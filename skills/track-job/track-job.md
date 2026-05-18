---
description: Record a job application after applying
---

Use this skill immediately after applying to any job.

You must make a curl request to the following URL using the rules below.

```bash
https://script.google.com/macros/s/AKfycbzOJEnyTAjXHhv_89Gm3hQLG_KwPHjxdYUIzB21ktuP2Vua0VJg-LKkRjPSua65AEA2Ng/exec
```

## Request
 
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body:** A JSON object matching the schema below
### Schema
 
| Field | Type | Required | Constraints |
|---|---|---|---|
| `link` | string | yes | URL of the job posting |
| `dateApplied` | string | yes | ISO 8601 date, e.g. `"2026-05-18"` |
| `industry` | string | yes | See valid values below |
| `phoneNumber` | string | no | 10–15 digits, common formatting characters allowed |
| `email` | string | yes | Employer contact email |
| `status` | string | yes | See valid values below |
| `notes` | string | no | Free text |
 
### Valid `status` Values
 
- `"Not Started"`
- `"Applied Only"`
- `"Applied + Emailed"`
- `"Applied + Called"`
- `"Applied + Emailed + Called"`
- `"Interview!"`
- `"Done"`
### Valid `industry` Values
 
- `"Tech"`
- `"Health Care"`
- `"Retail"`
- `"Finance"`
- `"Gig"`
- `"Other"`
## Example Request
 
```bash
curl -L -X POST "https://script.google.com/macros/s/AKfycbzOJEnyTAjXHhv_89Gm3hQLG_KwPHjxdYUIzB21ktuP2Vua0VJg-LKkRjPSua65AEA2Ng/exec" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com/job",
    "dateApplied": "2026-05-18",
    "industry": "Tech",
    "phoneNumber": "555-123-4567",
    "email": "hiring@example.com",
    "status": "Applied Only",
    "notes": "Referred by a friend"
  }'
```
 
> **Important:** The `-L` flag is required. Apps Script redirects before executing your function and curl will not follow the redirect without it.
