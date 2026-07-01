---
name: image
description: Use when the user attaches a screenshot/image
---

# Visual Code Mapping

The user can see more than you. If they have selected to provide an image to you - that means there is something significant in the image that they want you to see. The following workflow will tell you how to react when given an image. 

## Workflow

1. Analyze the image carefully. Identify the UI element, style issue, layout bug, feature request, or visual mismatch.
2. Search the codebase for the matching component, style, or logic. Verify likely matches with file names, labels, class names, accessibility labels, screenshots, or nearby text.
3. Map the gap between the image and current conversation context. Ask the user if intent is unclear.
4. Report:
   - Image shows
   - Found in
   - Current code
   - Problem
   - Proposed fix
5. Wait for confirmation before editing unless the user already asked for the fix.
