---
description: Record a job application after applying
---

Use this skill immediately after applying to any job.

Run the `track_job.py` script from the agents-data repo with the job details:

```bash
python /Users/farquaad/agents-data/skills/track-job/track_job.py "<Company>" "<Job Posting URL>" "Applied?" "<Comments>"
```

Parameters (in order):
**Name** — Company name and/or role title.
**Link** — URL of the job posting.
**Applied?** — Pass yes if you successfully submitted the application
**Comments** — Free-form notes on the job that are not implied or obvious. This is optional