<br />

<p align="center">
  <h3 align="center">
    Project Command Line Tool
</h3>
    <p align="center">
      A comfortable tools to start your projects!
      <br />
      <a href=""><strong>Explore the docs >></strong></a>
    </p>

</h3>
</p>

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->

## About The Project

This project is a tool for creating projects.



## Usage

This paragraph explains how to use this program.

### View project template list

```shell
pclt list
```
<img width="800" alt="pclt list" src="images/pclt-list.gif">

### Create

```shell
pclt create -pn <project-name>
```
<img width="800" alt="pclt create" src="images/pclt-create-normal.gif">

### Create (Other case)

* Spring Initializr

  ```shell
  pclt create -pn spring-init ./
  ```

  <img width="800" alt="pclt spring" src="images/pclt-spring-initializr.gif">

* Github

  ```shell
  pclt create-pn github ./
  ```
  <img width="800" alt="pclt github" src="images/pclt-create-github.gif">

### Delete

  ```shell
  pclt rm <project-name>
  ```
  OR
  ```shell
  pclt remove <project-name>
  ```

### Save

  ```shell
  pclt save
  ```