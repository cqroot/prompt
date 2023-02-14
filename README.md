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
- [Screenshots](#screenshots)
  - [Choose](#choose)
  - [MultiChoose](#multichoose)
  - [Input](#input)
  - [TextArea](#textarea)
  - [Toggle](#toggle)
  - [Show help message](#show-help-message)
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

<details><summary>View screenshots</summary>

**Theme Default**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose/main.go)

![choose](https://user-images.githubusercontent.com/46901748/218780062-a50a1043-f9ef-449a-8094-b412be831bfd.gif)

**Theme Arrow**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose-theme-arrow/main.go)

![choose-theme-arrow](https://user-images.githubusercontent.com/46901748/218780071-08b71a0b-963a-4078-9ac8-14ce89b02bd1.gif)

**Theme Line**

[example](https://github.com/cqroot/prompt/blob/main/examples/choose-theme-line/main.go)

![choose-theme-line](https://user-images.githubusercontent.com/46901748/218780074-c7988e70-a51d-4101-ae38-a2388989a269.gif)

</details>

### MultiChoose

<details><summary>View screenshots</summary>

**Theme Default**

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose/main.go)

![multichoose](https://user-images.githubusercontent.com/46901748/218780644-f32430fa-9d1e-4d8a-b197-8c9d89de9317.gif)

**Theme Dot**

[example](https://github.com/cqroot/prompt/blob/main/examples/multichoose-theme-dot/main.go)

![multichoose-theme-dot](https://user-images.githubusercontent.com/46901748/218780650-f3f216a9-7670-474b-8f5a-7b705a7ae1ab.gif)

</details>

### Input

<details><summary>View screenshots</summary>

[example](https://github.com/cqroot/prompt/blob/main/examples/input/main.go)

![screenshot-input](https://user-images.githubusercontent.com/46901748/216246350-d14074b0-0895-4a0b-890f-11c0cd725a04.gif)

**Password input**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-password/main.go)

![screenshot-input-echo-password](https://user-images.githubusercontent.com/46901748/218309893-754711a2-88f9-42de-a20f-6a86aeefeba0.gif)

**Password input like linux (do not display any characters)**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-echo-none/main.go)

![screenshot-input-echo-none](https://user-images.githubusercontent.com/46901748/218309957-468a1da3-f07c-4dc3-aa57-e1844f4f4c0e.gif)

**Only integers can be entered**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-integer-only/main.go)

**Only numbers can be entered**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-number-only/main.go)

**Input with validation**

[example](https://github.com/cqroot/prompt/blob/main/examples/input-with-validation/main.go)

![screenshot-input-with-validation](https://user-images.githubusercontent.com/46901748/218308650-dff43d9c-61d4-4373-8ac0-876ad2e329ae.gif)

</details>

### TextArea

<details><summary>View screenshots</summary>

[example](https://github.com/cqroot/prompt/blob/main/examples/textarea/main.go)

![screenshot-textarea](https://user-images.githubusercontent.com/46901748/218306061-d5f0ba9f-e6d5-43c8-ae04-88ae1cf8e758.gif)

</details>

### Toggle

Deprecated: use `Choose([]string{}, WithTheme(ChooseThemeLine))` instead.
[Preview](#theme-line).

### Show help message

<details><summary>View screenshots</summary>

All components support displaying help message for shortcut keys at the bottom.

![choose-with-help](https://user-images.githubusercontent.com/46901748/218780082-7808b54e-c258-427c-a91b-84b14ae7c246.gif)

Examples:

1. [Choose with help](https://github.com/cqroot/prompt/blob/main/examples/choose-with-help/main.go)
2. [MultiChoose with help](https://github.com/cqroot/prompt/blob/main/examples/multichoose-with-help/main.go)
3. [Input with help](https://github.com/cqroot/prompt/blob/main/examples/input-with-help/main.go)
4. [TextArea with help](https://github.com/cqroot/prompt/blob/main/examples/textarea-with-help/main.go)
5. [Toggle with help](https://github.com/cqroot/prompt/blob/main/examples/toggle-with-help/main.go)

![screenshot-help](https://user-images.githubusercontent.com/46901748/216308618-0b865448-23cd-4029-9a26-d6802b375fa4.png)

</details>

## License

[MIT License](https://github.com/cqroot/prompt/blob/main/LICENSE).
