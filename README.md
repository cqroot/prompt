<div align="center">
  <h1>Prompt</h1>
  <p>
    <i>User-friendly interactive prompts for Go.</i>
    <br />
    <i>
      Based on <a href="https://github.com/charmbracelet/bubbletea" alt="Bubble Tea">Bubble Tea</a>.
      Inspired by <a href="https://github.com/terkelg/prompts" alt="Prompts">Prompts</a>
        and <a href="https://github.com/charmbracelet/gum" alt="Gum">Gum</a>.
    </i>
  </p>

  <p>
    <a href="https://github.com/cqroot/prompt/actions">
      <img src="https://github.com/cqroot/prompt/workflows/test/badge.svg" alt="Action Status" />
    </a>
    <a href="https://codecov.io/gh/cqroot/prompt">
      <img src="https://codecov.io/gh/cqroot/prompt/branch/main/graph/badge.svg" alt="Codecov" />
    </a>
    <a href="https://goreportcard.com/report/github.com/cqroot/prompt">
      <img src="https://goreportcard.com/badge/github.com/cqroot/prompt" alt="Go Reference" />
    </a>
    <a href="https://pkg.go.dev/github.com/cqroot/prompt">
      <img src="https://pkg.go.dev/badge/github.com/cqroot/prompt.svg" alt="Go Reference" />
    </a>
    <a href="https://github.com/cqroot/prompt/tags">
      <img src="https://img.shields.io/github/v/tag/cqroot/prompt" alt="Git tag" />
    </a>
    <a href="https://github.com/cqroot/prompt/blob/main/go.mod">
      <img src="https://img.shields.io/github/go-mod/go-version/cqroot/prompt" alt="Go Version" />
    </a>
    <a href="https://github.com/cqroot/prompt/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/cqroot/prompt" />
    </a>
    <a href="https://github.com/cqroot/prompt/issues">
      <img src="https://img.shields.io/github/issues/cqroot/prompt" />
    </a>
  </p>
</div>

## Table of Contents

- [Features](#features)
- [Screenshots](#screenshots)
  - [Choose](#choose)
  - [MultiChoose](#multichoose)
  - [Input](#input)
  - [Write](#write)
- [License](#license)

## Features

1. `choose` lets the user choose one of several strings using the terminal ui.
2. `multichoose` lets the user choose multiple strings from multiple strings using the terminal ui.
3. `input` lets the user enter a string using the terminal ui.
   You can specify that only **numbers** or **integers** are allowed.
4. Show help message for keymaps.
5. Based on [Bubble Tea]("https://github.com/charmbracelet/bubbletea").
   `prompt.Prompt` and all child models implement `tea.Model`.

## Screenshots

### Choose

**Theme Default**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose/main.go)

![choose](https://user-images.githubusercontent.com/46901748/219288366-d4ce04df-ca98-4a03-8a80-e7c26577e86a.gif)

**Theme Arrow**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose-theme-arrow/main.go)

![choose-theme-arrow](https://user-images.githubusercontent.com/46901748/218780071-08b71a0b-963a-4078-9ac8-14ce89b02bd1.gif)

**Theme Line**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose-theme-line/main.go)

![choose-theme-line](https://user-images.githubusercontent.com/46901748/218780074-c7988e70-a51d-4101-ae38-a2388989a269.gif)

### MultiChoose

**Theme Default**

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose/main.go)

![multichoose](https://user-images.githubusercontent.com/46901748/219288777-1c913ac8-4144-4b96-b5be-3085483d8bae.gif)

**Theme Dot**

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose-theme-dot/main.go)

![multichoose-theme-dot](https://user-images.githubusercontent.com/46901748/218780650-f3f216a9-7670-474b-8f5a-7b705a7ae1ab.gif)

### Input

[example](https://github.com/cqroot/prompt/blob/main/examples/input/main.go)

![input](https://user-images.githubusercontent.com/46901748/219288988-12923602-a112-4876-906d-3575f3c50741.gif)

**Password input**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-password/main.go)

![input-echo-password](https://user-images.githubusercontent.com/46901748/218799172-ce501335-9821-4bf2-949a-0c08057d810f.gif)

**Password input like linux (do not display any characters)**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-none/main.go)

![input-echo-none](https://user-images.githubusercontent.com/46901748/218799167-59b52b0d-228e-4cb3-8bf2-7cf844874100.gif)

**Only integers can be entered**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-integer-only/main.go)

**Only numbers can be entered**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-number-only/main.go)

**Input with validation**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-with-validation/main.go)

![input-with-validation](https://user-images.githubusercontent.com/46901748/218799174-9355fcb1-bcef-4fe6-8421-e9472e913010.gif)

### Write

[example](https://github.com/cqroot/prompt/blob/main/examples/textarea/main.go)

![write](https://user-images.githubusercontent.com/46901748/219289253-7fef6708-c852-4d88-b2d0-376249f46c9b.gif)

## License

[MIT License](https://github.com/cqroot/prompt/blob/main/LICENSE).
