# Go-ExplainDT
A pure GO implementation of the work in "A Symbolic Language for Interpreting Decision Trees".

### Tests
To run test you must have docker installed and run the following command at
the root of the project:
```
docker build -f dockerfiles/Dockerfile.Tests -t goexpdt-tests --progress plain --no-cache-filter=run-tests-stage --target run-tests-stage .
```

### TODO
- Add correct simplification to circuits when passing `GuardedConst` as
  arguments.
- Look into passing `CNF` struct in `Encoding` methods to avoid creating to much
  garbage.
