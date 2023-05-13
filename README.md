# projector-assignment22

| operation                     | citus-vertical | ordinary |
|-------------------------------|----------------|----------|
| insert 1000000 rows           | 1m29.25s       | 1m14.69s |
| `SELECT avg(year) from books` | 120ms          | 80ms     |

## Setup:
   1. run `cd citus-vertical` or `cd ordinary` 
   2. run `docker-compose up`
   3. Execute table generation script in corresponding folder
   4. run inserter