<div align="center">
<img src="https://github.com/kartik-dutt/Personalized-Search-Engine/raw/main/images/SearchEngineLogo.png" height=150 width=150 alt="Logo">

<h1>Personalized-Search-Engine</h1>
<div align="center">
  <img src ="https://pkg.go.dev/static/img/badge.svg">
  <img src ="https://aleen42.github.io/badges/src/visual_studio_code.svg">
  <img src ="https://aleen42.github.io/badges/src/github.svg">
  <img src ="https://camo.githubusercontent.com/4e084bac046962268fcf7a8aaf3d4ac422d3327564f9685c9d1b57aa56b142e9/68747470733a2f2f7472617669732d63692e6f72672f6477796c2f657374612e7376673f6272616e63683d6d6173746572">
  </div>
<p>
Ever wanted to store the websites you frequently use? Store them using your personalized search engine. Search over million of URLs in couple of microseconds (us).
</p>

<hr/>
<h3>Demo</h3>

<img src="https://github.com/kartik-dutt/Personalized-Search-Engine/raw/main/images/SearchEngine.gif" alt="Demo">

<hr/>
<h3>About</h3>
<p>
This app provides a fast personalized search engine. Add documents that you often look up in the dataset or use the wiki-dataset that is provided. Search through millions of web pages in microseconds locally.

<b>Personalized Use Case :</b>

Add you stackoverflow links to the dataset and keep whenever you encounter the same problem again use this app, instead of spending hours on the net searching for the same link again.
</p>

<hr/>
<h3>Installation</h3>
Run the following commands to install the app.

```
git clone https://github.com/kartik-dutt/Personalized-Search-Engine.git
cd src
go mod tidy
go build
```

You may also need to download the dataset one time, using
```
cd data
go run download_dataset.go
```

<hr/>
<h3>Run the App</h3>
Run the executable using terminal or double clicking on the app.

```
./query
```
</div>

