[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]


<br />
<p align="center">

  <h1 align="center">nmapprom</h1>

  <p align="center">
    Prometheus exporter for counting the hosts connected to a network using nmap
    <br />
    ·
    <a href="https://github.com/OisinA/nmapprom/issues">Report Bug</a>
    ·
    <a href="https://github.com/OisinA/nmapprom/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

nmapprom is a Prometheus exporter for counting the hosts connected to a network using nmap, written in Go.



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* nmap
* Go

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/OisinA/nmapprom.git
   ```
2. Build the go project
   ```sh
   go build
   ```

<!-- USAGE EXAMPLES -->
## Usage

nmapprom requires passing in the ip range of your local network, checking to see what addresses are in use.

```
nmapprom -i 192.168.0.0/24
```


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/OisinA/nmapprom.svg?style=for-the-badge
[contributors-url]: https://github.com/OisinA/nmapprom/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/OisinA/nmapprom.svg?style=for-the-badge
[forks-url]: https://github.com/OisinA/nmapprom/network/members
[stars-shield]: https://img.shields.io/github/stars/OisinA/nmapprom.svg?style=for-the-badge
[stars-url]: https://github.com/OisinA/nmapprom/stargazers
[issues-shield]: https://img.shields.io/github/issues/OisinA/nmapprom.svg?style=for-the-badge
[issues-url]: https://github.com/OisinA/nmapprom/issues
[license-shield]: https://img.shields.io/github/license/OisinA/nmapprom.svg?style=for-the-badge
[license-url]: https://github.com/OisinA/nmapprom/blob/master/LICENSE.txt