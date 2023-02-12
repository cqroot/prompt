<div align="center">
  <h1>Prompt</h1>
  <p>
    <i>User-friendly interactive prompts for Go.</i>
    <br />
    <i>
      Based on <a href="https://github.com/charmbracelet/bubbletea" alt="Bubble Tea">Bubble Tea</a>.
      Inspired by <a href="https://github.com/terkelg/prompts" alt="prompts">Prompts</a>
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
- [Usage](#usage)
  - [Input](#input)
    - [Password input](#password-input)
    - [Password input like linux (do not display any characters)](<#password-input-like-linux-(do-not-display-any-characters)>)
    - [Only integers can be entered](#only-integers-can-be-entered)
    - [Only numbers can be entered](#only-numbers-can-be-entered)
    - [Input with validation](#input-with-validation)
  - [TextArea](#textarea)
  - [Toggle](#toggle)
  - [Choose](#choose)
  - [MultiChoose](#multichoose)
  - [Show help message](#show-help-message)
- [License](#license)

## Features

1. `input` lets the user enter a string using the terminal ui.
   You can specify that only **numbers** or **integers** are allowed.
2. `toggle` lets the user choose one of several strings using the terminal ui (Usually used for yes or no choices).
3. `choose` lets the user choose one of several strings using the terminal ui.
4. `multichoose` lets the user choose multiple strings from multiple strings using the terminal ui.
5. Show help message for keymaps.
6. Based on [Bubble Tea]("https://github.com/charmbracelet/bubbletea").
   `prompt.Prompt` and all child models implement `tea.Model`.

## Usage

### Input

[example](https://github.com/cqroot/prompt/blob/main/examples/input/main.go)

![screenshot-input](https://user-images.githubusercontent.com/46901748/216246350-d14074b0-0895-4a0b-890f-11c0cd725a04.gif)

#### Password input

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-password/main.go)

![screenshot-input-echo-password](https://user-images.githubusercontent.com/46901748/218309893-754711a2-88f9-42de-a20f-6a86aeefeba0.gif)

#### Password input like linux (do not display any characters)

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-none/main.go)

![screenshot-input-echo-none](https://user-images.githubusercontent.com/46901748/218309957-468a1da3-f07c-4dc3-aa57-e1844f4f4c0e.gif)

#### Only integers can be entered

[example](https://github.com/cqroot/prompt/blob/main/examples/input-integer-only/main.go)

#### Only numbers can be entered

[example](https://github.com/cqroot/prompt/blob/main/examples/input-number-only/main.go)

#### Input with validation

[example](https://github.com/cqroot/prompt/blob/main/examples/input-with-validation/main.go)

![screenshot-input-with-validation](https://user-images.githubusercontent.com/46901748/218308650-dff43d9c-61d4-4373-8ac0-876ad2e329ae.gif)

### TextArea

[example](https://github.com/cqroot/prompt/blob/main/examples/textarea/main.go)

![screenshot-textarea](https://user-images.githubusercontent.com/46901748/218306061-d5f0ba9f-e6d5-43c8-ae04-88ae1cf8e758.gif)

### Toggle

[example](https://github.com/cqroot/prompt/blob/main/examples/toggle/main.go)

![screenshot-toggle](https://user-images.githubusercontent.com/46901748/216246356-fb3eb7df-7240-4a09-8899-45797bfe79c7.gif)

### Choose

[example](https://github.com/cqroot/prompt/blob/main/examples/choose/main.go)

![screenshot-choose](https://user-images.githubusercontent.com/46901748/216246342-da8d8b67-983c-41b8-b85d-a4ef2dcab0bd.gif)

### MultiChoose

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose/main.go)

![screenshot-multichoose](https://user-images.githubusercontent.com/46901748/216246355-92129b7b-c812-4b15-bfbc-7ec7e39e972a.gif)

### Show help message

`Prompt.WithHelp(true)` displays the help message for keymaps.

```go
val, err := prompt.New().Ask("Choose value:").WithHelp(true).
	Choose([]string{"Item 1", "Item 2", "Item 3"})
```

Examples:

1. [Input with help](https://github.com/cqroot/prompt/blob/main/examples/input-with-help/main.go)
2. [TextArea with help](https://github.com/cqroot/prompt/blob/main/examples/textarea-with-help/main.go)
3. [Toggle with help](https://github.com/cqroot/prompt/blob/main/examples/toggle-with-help/main.go)
4. [Choose with help](https://github.com/cqroot/prompt/blob/main/examples/choose-with-help/main.go)
5. [MultiChoose with help](https://github.com/cqroot/prompt/blob/main/examples/multichoose-with-help/main.go)

![screenshot-help](https://user-images.githubusercontent.com/46901748/216308618-0b865448-23cd-4029-9a26-d6802b375fa4.png)

## License

[MIT License](https://github.com/cqroot/prompt/blob/main/LICENSE).
