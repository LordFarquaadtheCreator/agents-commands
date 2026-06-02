---
name: record-a-job-application
description: you must use the skill after applying to a job
metadata:
  display-name: Record a Job Application
  enabled: 'true'
---
# description: Record a job application after applying

Use this skill immediately after applying to any job. If the command fails, you must stop and tell the user that the command failed. A failed command exits with code 1 with clear error formatting.

Use the `track-job.py` script to record job applications:

```zsh
/Users/farquaad/.pyenv/shims/python3 /Users/farquaad/agents-data/skills/track-job/track-job.py <Job Posting URL> <industry> <status> [email] [phone] [notes]
```

Parameters (in order):
- **Link** — URL of the job posting (this must be the link relating to the individual job application, this link is always attainable by searching through share buttons and copy link buttons)
- **Industry** — values in teh set of: Tech, Health Care, Retail, Finance, Gig, Other.
- **Status** — values in the set of: Not Started, Applied Only, Applied + Emailed, Applied + Called, Applied + Emailed + Called, Interview!, Done.
- **Email** — Employer contact email (optional)
- **Phone** — Contact phone number (optional).
- **Notes** — Free-form notes on the job (optional you do not need to fill this out if there is nothing special to remark about the job).

**Optional parameters can be omitted from the function call**

## Example

*No optional parameters passed*
```zsh
/Users/farquaad/.pyenv/shims/python3 /Users/farquaad/agents-data/skills/track-job/track-job.py "https://fakejobs.com/quantum-ai-analyst" "Tech" "Not Started"
```

*Some optional parameters passed*
```zsh
/Users/farquaad/.pyenv/shims/python3 /Users/farquaad/agents-data/skills/track-job/track-job.py "https://fakejobs.com/quantum-ai-analyst" "Tech" "Not Started" "email@email.com"
```

*All optional parameters passed*
```zsh
/Users/farquaad/.pyenv/shims/python3 /Users/farquaad/agents-data/skills/track-job/track-job.py "https://fakejobs.com/quantum-ai-analyst" "Tech" "Not Started" "email@email.com" "917-999-1234" "They said to email \"John\" at \"john@company.com\""
```