# RepunKamuy
iproduction application dispatch, render and computing for Docker.

https://docs.docker.com/develop/sdk/ 
https://github.com/mottosso/cmdx#third-party
https://github.com/kubernetes/client-go

## Why
開発の終了したAtheraでフルパッケージ提供されていたものを分解して、
スタジオユースに適した環境構築、試行錯誤に適したtuiベースのオーケストレーションツールです。

ツールのテストや時間のかかるディスパッチャを気軽に試す環境を準備することは、
時間的制約や人的コストの観点から大変です。
スタジオユースなコンテナはありますが 個人で使えるイメージを必要になったときに
使うことはインフラストラクチャ部門との調整が必要になり難しいです。

アウトソース委託先の環境を想定したライブラリ制作
バージョンの異なるアプリケーション間での動作保証

RepunKamuy はローカルマシン1台から始める事のできるVFXプラットフォームの土台です。
GCP AWS Azure を組み合わせたハイブリッド環境を構築する手助けをします。
ZYNC, OpenCue, Teradici と組み合わせたオンプレミスとのハイブリッド環境を
構築することも出来ます。

初期サポートではBlender, Houdini, UnrealEngine4 をサポートします。

コンテナイメージから起動する際にシェルベースでセットアップを行うのは大変


UnrealGCP
Blender Container


### api features
- シンプルなコード(set, run, end)
- Python API の提供(via grumpy)
- VFX platform　をベースにした構成

```go
{ "apps": enum, "packages": enum, "os": enum, "foundation": bool, "Num": int, "gpu": bool, "licenses": [string], "script": string, "file": file "plugins": [string] }

enum apps { blender, houdini, unreal }

enum os {ubuntu, centos, win10 }

enum lib { gcc,cuda,}

enum packages { fbx, usd, ffmpeg, opencv, arnold, vray }
