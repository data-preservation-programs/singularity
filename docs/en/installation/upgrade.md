# Version Upgrade

Between each minor version upgrade, e.g. 2.3.0 to 2.4.0, use below command to upgrade the database schema:
```bash
singularity admin init
```

For major version upgrade, e.g. 2.4.0 to 3.0.0, please refer to migration guide for each major version release.

If you are moving existing schedules from legacy market deals to PDP or DDO, see [Migrating from legacy deals](migrating-from-legacy-deals.md).
