# podscript

podscript is a tool to generate transcripts for podcasts (and other similar audio files), using LLMs and other Speech-to-Text (STT) APIs.

## Install

```shell
> go install github.com/deepakjois/podscript@latest

> ~/go/bin/podscript --help
```

## Configure

This command displays prompts to enter API keys for supported services, and write them to `$HOME/.podscript.toml`.

```shell
> podscript configure
```

Alternatively, you can set keys in environment variables – `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, `GROQ_API_KEY`, `DEEPGRAM_API_KEY`

## Transcript from YouTube videos

For podcasts on YouTube with autogenerated captions (e.g. [Andrew Huberman](https://www.youtube.com/watch?v=WFcYF_pxLgA) and [Cal Newport](https://www.youtube.com/watch?v=OvlfCW3Ec1g)), use the `ytt` subcommand to download the captions from the YouTube video and feed it to an LLM model to generate a clean transcript. 

```shell
> podscript ytt https://www.youtube.com/watch?v=aO1-6X_f74M
```

It uses the `gpt-4o` model by default. Use `--model` flag to set a different model. The following are supported:

- `gpt-4o`
- `gpt-4o-mini`
- `claude-3-5-sonnet-20241022`
- `claude-3-5-haiku-20241022`
- `llama-3.3-70b-versatile`
- `llama-3.1-8b-instant`

## Transcript from audio URLs and files
If the podcast (or other audio) is not available on YouTube, or you need higher quality transcripts, you can send the audio to a Speech-to-Text (STT) API.

podscript supports the following STT APIs:

- [Deepgram](https://playground.deepgram.com/?endpoint=listen&smart_format=true&language=en&model=nova-2) (which as of Jan 2025 provides $200 free signup credit!)
- [Groq](https://console.groq.com/docs/speech-text) (which as of Jul 2024 is in beta and free to use within your rate limits).
- [Assembly AI](https://www.assemblyai.com/docs) (which as of Oct 2024 is free to use within your credit limits and they provide $50 credits free on signup).

> [!TIP]
> You can find the audio download link for a podcast on [ListenNotes](https://www.listennotes.com/) under the More menu
>
> <img width="252" alt="image" src="https://github.com/deepakjois/podscript/assets/5342/1f400964-e575-4f59-9de0-ee75f386b27d">


### Example Usage

Deepgram and AssemblyAI subcommands support `--from-url` for passing audio URLs, and `--from-file` to pass audio files. Groq only supports audio files.

All the subcommands support the `-o` flag to write the output to a text file. Other options to set the model, or dump the full API response are provided where available.

#### Deepgram

```shell
> podscript deepgram --from-url  https://audio.listennotes.com/e/p/d6cc86364eb540c1a30a1cac2b77b82c/
```

#### Groq

```shell
> podscript groq --file huberman.mp3
```

#### AssemblyAI

```shell
> podscript assembly-ai --from-url https://audio.listennotes.com/e/p/d6cc86364eb540c1a30a1cac2b77b82c/ -o transcript.txt
```

## Feedback

Feel free to drop me a note on [X](https://x.com/debugjois) or [Email Me](mailto:deepak.jois@gmail.com)

## License

[MIT](https://github.com/deepakjois/podscript/raw/main/LICENSE)
