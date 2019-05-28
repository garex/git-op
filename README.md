<img src="https://user-images.githubusercontent.com/77981/58380091-cec11500-7fd6-11e9-8325-c2bbcdab2cc8.png" alt="git op logo" align="left" width="120" height="120" />

# git-op
Git branching tool.

One line table of contents: [Installing](#installing) | Usage: [Start - Branch](#start---branch), [Finish - Merge](#finish---merge), [Release - Tag](#release---tag).

## Installing

Download your's OS/architecture binary into your $PATH from [latest release page](https://github.com/garex/git-op/releases/latest).

For linux/amd64:

```bash
sudo wget -O /usr/local/bin/git-op https://github.com/garex/git-op/releases/download/1.1/git-op
sudo chmod +x /usr/local/bin/git-op
```

## Usage

### Start - Branch

`git op branch [NAME]` creates branch from 'master' named 'NAME'

When `[NAME]` is ommitted, it autogenerated like '2019-05-26-13-09'.

### Finish - Merge

`git op merge [NAME]` merges branch 'NAME' into 'master'.

When `[NAME]` is ommitted, current branch is used. On 'master' branch it do nothing.

Default behavior pulls latest 'master' branch from 'origin' remote and rebases onto it. Tags are fetched too.

### Release - Tag

`git op tag [VERSION]` creates version as a tag.

Assume we have last tag '1.2.3'.

* `[VERSION]` is ommited or 'patch' passed -- creates next 'patch' version: '1.2.4'.
* 'minor' passed -- '1.3'.
* 'major' passed -- '2.0'.

Default behavior calls `git op merge` before releasing.
