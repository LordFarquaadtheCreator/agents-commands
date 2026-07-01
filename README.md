# agents-skills

Canonical skill bank for Zed agent skills. Each skill in `skills/` is the source of truth.

## Adding a skill to an app

Skills are **not** auto-discovered. You must manually symlink each skill into an app's agent skill bank:

```zsh
ln -s "$PWD/skills/<skill-name>" /path/to/app/skills/<skill-name>
```

This is intentional — you control exactly which skills each app can use. Two apps can have different sets of skills without one polluting the other.

Also, each skill is now able to be version tracked and managed through one source while possibly being available to any number of agents and skill banks.
