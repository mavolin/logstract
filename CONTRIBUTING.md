# Contributing

We would love to see the ideas you want to bring in to improve this project.
But before you get started, make sure to follow the guidelines below:

## Contributing through issues

If have an idea how to improve this project or if you found a bug, let us know by submitting an issue.
The issue templates will take care of most of the requirements, but there is one thing you should note:

### Titles

We not only use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for commits, but also for issue titles.
If you propose a feature, use `feat(PACKAGE_NAME): TITLE`, for a bug replace the `feat` with `fix`, for `docs` use `docs`, you get the hang.
Other types are `style` `refactor` (e.g. for shortening code) and `test`.
If your change is breaking (in the semantic versioning sense) add an exclamation mark behind the scope, e.g. `feat(package)!: title`.



## Code Contributions
### Opening an Issue

Before you hand in a PR, open an issue and tell us you'll be handling in an PR.
This gives us the ability to point out important things you should keep in mind and to tell you if such a feature would fit into this project at all.

### Commiting

This is by far the most important guideline.
Please make small, thoughtfull commits, a commit like `feat: add xy` with 20 new files is rarely appropriate.

#### Conventional Commits

Please use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for your contributions, once you get the hang of it, you'll see that they are pretty easy to use.
Just have a look at the [quick start guide](https://www.conventionalcommits.org/en/v1.0.0/#summary) on their website.
The scope is typically the package name, but for non-go files appropriate scopes may also be: `git`, `README` or `go.mod`.
If none match what you are doing, just think of something.
The types we use are: `fix`, `feat`, `docs`, `style`, `refactor` and `test`.
Breaking changes should be signaled using a `!`, and not by the footer.

### Fixing a Bug

If you're fixing a bug, make sure to add a test case for that bug, to make sure it's gone for good.
This of course only applies if the function is testable.

### Code Style

Make sure all code is `gofmt -s`'ed and passes the golangci-lint checks.
If you're code fails a lint task, but the way you did it is justified, add an execption to the `.golangci.yml` file with a comment explaining, why this exception exists.

### Testing

If possible and appropriate you should fully test the code you submit.
Each function exposed and unexposed should have a signle test, which either tests directly or is split into subtests, preferrably table-driven.

#### Naming

Please prefix your variables correctly:
* `testX` for test data that is passed to the tested function.
  These variables are declared first.
* `expect` (or `expectX` for multiple return values, errors excluded) is the value expected to be returned from the tested function.
  These variable come second.
* `actual` (or `actualX` for multiple return values, errors excluded) is the value actually returned by the tested function.
  These variable come last.

#### Error Reports

If `actual != expect` or `!reflect.DeepEqual(actual, expect)`, an error in the form of: `expected #{name} to return: #{expect}, but got #{actual}` should be reported (`t.Errorf` or `t.Fatalf`).
Name is either `package.struct.function` or `package.function` and `actual` and `expect` are formatted using `%+v`.
If a function unexpectedly returns an error, the error message should be formatted like so: `#{name} returned an error: #{err}` where name uses the same scheme as above and err is the value of `err.Error()`.

#### Table-Driven Tests

If there is a single table, it should be called testCases, multiple use the name `<type>Cases`, e.g. `successCases` and `failureCases`.
A table of length 1 may be replaced by a standard `t.Run`.

### Opening a Pull Request

When opening a pull request, merge against `develop` and use the title of the issue as the PR title.

A Pull Request must pass all test, before being merged.
