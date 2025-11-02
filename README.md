<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Hassembler</h3>

  <p align="center">
    A blazing-fast assembler for the Hack machine language.
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

This project implements a functional assembler for the Hack machine language. Hack is a fully functional machine language coding the Hack computer. For full specification check <a href="[#readme-top](https://www.nand2tetris.org/software)">Nand2Tetris</a> website.



### Built With

[![Golang][Golang]][Golang-url]




<!-- GETTING STARTED -->
## Getting Started

Purpose of this assembler is being as simple as possible while implementing all features of the Hack specification.

### Prerequisites

This project uses only the Go standard library. Hence, a working Go installation is all you need.

* For Ubuntu:
  ```sh
  sudo apt update
  sudo apt install go
  ```

### Building

* While in the project root, run:
  ```sh
  make
  ```

Built binary is placed inside the project root.

### Installation

Generated binary is statically linked, so you can move it anywhere as you like.

* To make it available in your shell:
```sh
  mv ./hassembler /usr/bin
```




<!-- USAGE EXAMPLES -->
## Usage

Hassembler consume only the first file given as command line argument.

* Example usage:
```sh
  hassembler l33t_game.asm
```

After a successful run, it places the generated machine code to current working directory.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Golang]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[Golang-url]: https://go.dev/


<p align="right">(<a href="#readme-top">back to top</a>)</p>
