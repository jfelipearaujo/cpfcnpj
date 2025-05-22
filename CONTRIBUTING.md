# Contributing to cpfcnpj

When contributing to this repository, please first discuss the change you wish to make via issue with the owners of this repository before making a change. This will ensure that your change is aligned with the goals of the project and that it can be reviewed and merged more easily.

Please note we have a code of conduct, please follow it in all your interactions with the project.

## Table of Contents

- [Contributing to cpfcnpj](#contributing-to-cpfcnpj)
  - [Table of Contents](#table-of-contents)
  - [Cloning the repository](#cloning-the-repository)
  - [Pull Request Process](#pull-request-process)
    - [Your PR is merged](#your-pr-is-merged)
    - [Release of a new version](#release-of-a-new-version)
  - [Code of Conduct](#code-of-conduct)

## Cloning the repository

```bash
git clone https://github.com/jfelipearaujo/cpfcnpj.git
```

## Pull Request Process

1. Ensure that you have the latest version of the codebase.
2. Configure your machine to use the `.pre-commit-config.yaml` file. **THIS IS IMPORTANT!**
   - Install pre-commit via Python (PIP): `pip install pre-commit`
   - Run `pre-commit install` to install the pre-commit hooks.
3. Create a new branch from the `main` branch.
4. Make your changes.
5. Ensure that your code passes all tests and linting checks.
6. Commit your changes with a clear and descriptive message.
7. Push your changes to your forked repository.
8. Create a new pull request from your branch to the `main` branch.

### Your PR is merged

Once your PR is merged, you will be able to see your changes in the `main` branch and you will be publicly visible on the repository contributors list.

### Release of a new version

When a new version is ready to be released, the maintainers will generate a new release and everyone who contributed to the repository will be mentioned in the release notes.

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.
