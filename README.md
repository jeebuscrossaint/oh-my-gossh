# oh-my-gossh

The most customizable and feature rich portfolio TUI for you. Written using the [Bubble Tea framework](https://github.com/charmbracelet/bubbletea) and the power of Golang. Letting you express your portfolio in an incredibly distinct way, utilizing the power of the TOML configuration file. Perhaps it may even make viewers say ***"Oh my gossh!"***.

____

### Configuration
A full configuration guide with an image example can be found [here](configs/configs.md).
Extremely modular configuration, core configuration should take no more than a few minutes. 

### Hosting
You can host it through docker, or just run it locally.

### Installation

1. **First, clone the repository and cd in.**

```bash
git clone github.com/jeebuscrossaint/oh-my-gossh
cd oh-my-gossh
```

2. **Next, in the folder, run the following commands to build the binary. (requires golang installed of course along with make)**

```bash
make release
make install
make firsttime
```
3. **Read configs/configs.md and configure your portfolio colors and TOML and make your markdown.**

4. **Run the binary.**

```bash
oh-my-gossh
```
