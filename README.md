# 🐳 Docker Color Output 2.x

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
<a href="https://go.dev"><img src="https://img.shields.io/github/go-mod/go-version/devemio/docker-color-output" alt="Go"></a>
<a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/v/release/devemio/docker-color-output" alt="Release"></a>
<a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/downloads/devemio/docker-color-output/total" alt="Downloads"></a>
</p>


## 👨‍💻 Installation

This package uses pipeline and allows you to color the Docker output.

#### 🍏 Mac

```bash
brew install dldash/core/docker-color-output
```

#### 🏝 Linux

```bash
sudo add-apt-repository ppa:dldash/core
sudo apt update
sudo apt install docker-color-output
```

#### 🪄 Aliases

For convenience use these [aliases](bash/aliases.sh).

## 📚 Usage

#### 💡 docker images

```bash
di # alias
```

```bash
docker images | docker-color-output
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

#### 💡 docker ps

```bash
dps [-a] # alias
```

```bash
docker ps [-a] | docker-color-output
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

#### 💡 docker compose ps

⚠️ The latest version works with docker-compose `2.x`.

```bash
dcps # alias
```

```bash
docker compose ps | docker-color-output
```

![docker compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)
