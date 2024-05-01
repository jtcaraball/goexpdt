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

The experiment outputs will be written to `cmd/experiment/output` directory.

### Commands

- `list`: List all implemented experiments.
- `info <experiment>`: Get experiment info and expected arguments.
- `<experiment> <args>`: Run experiment with arguments.

### Input Types

Experiments may accept one of two file formats as types, both of which must
be in the `cmd/experiment/input` directory.

- **Tree file**: A json file representing a decision tree.
- **Optimization file**: A plain text file that must follow the format outlined
  bellow

  ```
  <tree_file_name>
  <instance_1>
  <instance_2>
  ...
  <instance_n>
  ```

  Here `<tree_file_name>` corresponds to the name of a Tree file in the input
  directory and `<instance_i>` to an instance represented as a word in the
  alphabet {0, 1, 2} with 2 meaning that a feature is a 'bottom'.

### Examples

In the `cmd/experiments/inputs` directory there are examples of
tree and optimization file inputs. Here are some of the experiments that
can be ran on this inputs:

**Times for 10 random positive instances for optimal Sufficient Reason over
Less Level order**:

```
docker run --rm \
    -v $(pwd)/cmd/experiment/outputs:/goexpdt/cmd/experiment/outputs \
    -v $(pwd)/cmd/experiment/inputs:/goexpdt/cmd/experiment/inputs \
    goexpdt-exp optim:rand:stats:sr-ll 10 mnist_d0_n400.json
```

**Values of optimal Changed Allowed over Greater Hamming distance order for
specific instances**:

```
docker run --rm \
    -v $(pwd)/cmd/experiment/outputs:/goexpdt/cmd/experiment/outputs \
    -v $(pwd)/cmd/experiment/inputs:/goexpdt/cmd/experiment/inputs \
    goexpdt-exp optim:val:ca-gh mnist_d0_input.txt
```

## TODO

- Add naming convention mechanism for context variables.
- Use constants `BotCount` method in circuits.
- Add correct simplification to circuits when passing `GuardedConst` as
  arguments.
- Look into passing `CNF` struct in `Encoding` methods to avoid creating to much
  garbage.
