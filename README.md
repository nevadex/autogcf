# MAD BECAUSE "." ISNT SUPPORTED ON GCF URLS
### This prevents an index.html so defeats the purpose of this tool, which was finished by the time I realized
This is likely because GCF's backend is Cloud Run, which uses a subdomain as opposed to a path to host the function. (Meaning that GCF is not native functions, rather a wrapper!!!!).

Lucky for me, I spent 2 whole days on this. Not sure if that's a lot or not but I'm sure I could've popped off on siege I guess

---

# AutoGCF
### Automatically serve files from Google Cloud Functions (FOR FREE*)
This tool can be ran in Google Cloud Shell to make it easier

FREE* only if your file sizes and traffic is reasonable (<5gb bandwidth, <2mil visits per month)

## Usage:
- Install this with `go get github.com/nevadex/autogcf`
- Ensure gcloud is set up, and the correct account and project is set
- Run `autogcf -h` to see all options

## Limitations:
- Root-level files in your website source cannot have extensions. However, files in directories CAN have extensions.

### The best usage of this tool is to host your static assets, even regionally, like a cdn type of thing.
I highly recommend you read Google's docs on GCF before using this. It really could be a good, free CDN for your shiz.