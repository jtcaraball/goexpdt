# Go-ExplainDT

A pure GO implementation of the work in "A Symbolic Language for Interpreting Decision Trees".

## Tests

To run test you must have docker installed and run the following command at
the root of the project:

```
docker build -f dockerfiles/Dockerfile.Tests -t goexpdt-tests --progress plain --no-cache-filter=run-tests-stage --target run-tests-stage .
```

## Experiments

To build the experiment image you must have docker installed and run the
following command at the root of the project:

```
docker build -f dockerfiles/Dockerfile.Experiments -t goexpdt-exp .
```

After this to run tests run the following command:

```
docker run --rm \
    -v $(pwd)/cmd/experiment/outputs:/goexpdt/cmd/experiment/outputs \
    -v $(pwd)/cmd/experiment/inputs:/goexpdt/cmd/experiment/inputs \
    goexpdt-exp <command> <args>
```

The experiment outputs will be written to `cmd/experiment/output` dir.

### Commands

- `list`: List all implemented experiments.
- `info <experiment>`: Get experiment info and expected arguments.
- `<experiment> <args>`: Run experiment with arguments.

## TODO

- Add naming convention mechanism for context variables.
- Use constants `BotCount` method in circuits.
- Add correct simplification to circuits when passing `GuardedConst` as
  arguments.
- Look into passing `CNF` struct in `Encoding` methods to avoid creating to much
  garbage.
