# Contributing to MpesaOverlay

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## We Develop with Github

We use github to host code, to track issues and feature requests, as well as accept pull requests.

## Reporting issues

Reporting issues are a great way to contribute to the project. We are perpetually grateful about a well-written, thorough bug report.

Before raising a new issue, check [our issue list](https://github.com/0x6flab/mpesaoverlay/issues) to determine if it already contains the problem that you are facing.

A good bug report shouldn't leave others needing to chase you for more information. Please be as detailed as possible.

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

People _love_ thorough bug reports. I'm not even kidding.

## Pull requests

Good pull requests (e.g. patches, improvements, new features) are a fantastic help. They should remain focused in scope and avoid unrelated commits.

**Please ask first** before embarking on any significant pull request (e.g. implementing new features, refactoring code etc.), otherwise you risk spending a lot of time working on something that the maintainers might not want to merge into the project.

Please adhere to the coding conventions used throughout the project. If in doubt, consult the [Effective Go](https://golang.org/doc/effective_go.html) style guide.

To contribute to the project, [fork](https://help.github.com/articles/fork-a-repo/) it, clone your fork repository, and configure the remotes:

```bash
git clone https://github.com/<your-username>/mpesaoverlay.git
cd mpesaoverlay
git remote add upstream https://github.com/0x6flab/mpesaoverlay.git
```

If your cloned repository is behind the upstream commits, then get the latest changes from upstream:

```bash
git checkout main
git pull --rebase upstream main
```

Create a new topic branch from `main` using the naming convention `MO-[issue-number]` to help us keep track of your contribution scope:

```bash
git checkout -b MO-[issue-number]
```

Commit your changes in logical chunks. When you are ready to commit, make sure to write a Good Commit Message. Consult the [Conventional Commits guide](https://www.conventionalcommits.org/en/v1.0.0/) if you're unsure of what constitutes a Good Commit Message. Use [commitgpt](https://pypi.org/project/commitgpt/) to generate a Good Commit Message.

Note that every commit you make must be signed. By signing off your work you indicate that you are accepting the [Developer Certificate of Origin](https://developercertificate.org/).

Use your real name (sorry, no pseudonyms or anonymous contributions). If you set your `user.name` and `user.email` git configs, you can sign your commit automatically with `git commit -s`.

Locally merge (or rebase) the upstream development branch into your topic branch:

```bash
git pull --rebase upstream main
```

Push your topic branch up to your fork:

```bash
git push origin MO-[issue-number]
```

[Open a Pull Request](https://help.github.com/articles/using-pull-requests/) with a clear title and detailed description.

## License

By contributing, you agree that your contributions will be licensed under [Apache License](https://github.com/0x6flab/mpesaoverlay/blob/main/LICENSE) that covers the project.
