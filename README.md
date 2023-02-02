<div align="center">
  <h1>Prompt</h1>
  <p><i>User-friendly interactive prompts for Go.</i></p>

  <p>
    <a href="https://github.com/cqroot/prompt/actions">
      <img src="https://github.com/cqroot/prompt/workflows/test/badge.svg" alt="Action Status" />
    </a>
    <a href="https://pkg.go.dev/github.com/cqroot/prompt" target="_blank">
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

# Features

1. `input` lets the user enter a string using the terminal ui.
2. `toggle` lets the user choose one of several strings using the terminal ui (Usually used for yes or no choices).
3. `choose` lets the user choose one of several strings using the terminal ui.
4. `multichoose` lets the user choose multiple strings from multiple strings using the terminal ui.

# Usage

## Input

[example](https://github.com/cqroot/prompt/blob/main/examples/input/main.go)

![screenshot-input](https://user-images.githubusercontent.com/46901748/216246350-d14074b0-0895-4a0b-890f-11c0cd725a04.gif)

`InputWithLimit` can limit the type of input:

```go
// Only integers can be entered
val, err = p.Ask("Input example (Only Integer):").InputWithLimit("", prompt.InputInteger)

// Only numbers (integers and decimals) can be entered
val, err = p.Ask("Input example (Only Number):").InputWithLimit("", prompt.InputNumber)
```

## Toggle

[example](https://github.com/cqroot/prompt/blob/main/examples/toggle/main.go)

![screenshot-toggle](https://user-images.githubusercontent.com/46901748/216246356-fb3eb7df-7240-4a09-8899-45797bfe79c7.gif)

## Choose

[example](https://github.com/cqroot/prompt/blob/main/examples/choose/main.go)

![screenshot-choose](https://user-images.githubusercontent.com/46901748/216246342-da8d8b67-983c-41b8-b85d-a4ef2dcab0bd.gif)

## MultiChoose

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose/main.go)

![screenshot-multichoose](https://user-images.githubusercontent.com/46901748/216246355-92129b7b-c812-4b15-bfbc-7ec7e39e972a.gif)

## Show help message

`Prompt.SetHelpVisible(true)` displays the help message for key bindings.

```go
val, err := prompt.New().Ask("Choose value:").SetHelpVisible(true).
	Choose([]string{"Item 1", "Item 2", "Item 3"})
```

![screenshot-help](https://user-images.githubusercontent.com/46901748/216308618-0b865448-23cd-4029-9a26-d6802b375fa4.png)

# License

Released under the [MIT License](https://github.com/cqroot/prompt/blob/main/LICENSE).
