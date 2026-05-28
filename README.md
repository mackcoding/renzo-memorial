# Renzo Memorial Website

This is a GitHub-hosted memorial website for our gentle giant, Renzo.

## 🐕 About Renzo

Renzo was a loyal, patient German Shepherd with a heart as big as he was. He spent most of his life beside his bossy yellow Labrador brother, Elliott, taking the stolen toys and stolen sunbeams in stride with the unshakable patience of a true gentleman.

Renzo passed away on May 10, 2026, and this website serves as a digital memorial to celebrate his life and preserve precious memories. His full obituary lives at [hollowaypets.com](https://www.hollowaypets.com/obituaries/Renzo?obId=48365295).

## 🏗️ How It Works

This memorial website is built as a static site hosted on GitHub Pages with an automated photo gallery system:

### Architecture

- **Frontend**: Simple HTML/CSS/JavaScript static site using Bootstrap for responsive design
- **Photo Gallery**: Dynamic JSON-based gallery system that automatically indexes photos
- **Backend**: Go-based photo indexer that processes images and generates gallery data
- **Deployment**: GitHub Actions workflow for automated building and deployment

### Photo Gallery System

The heart of this site is the automated photo gallery powered by a Go script (`scripts/indexer.go`):

1. **Photo Scanning**: Scans the `web/photos/` directory for image files
2. **Description Loading**: Reads descriptions from `web/photos/description.json`
3. **Pagination**: Automatically creates paginated gallery data (15 photos per page)
4. **JSON Generation**: Creates `index.json` and `pageN.json` files in `web/index/`
5. **Frontend Integration**: JavaScript loads and displays the gallery dynamically

### Automated Workflow

The GitHub Actions workflow (`deploy.yml`) handles everything automatically:

1. **Triggers**: Runs on any push to the main branch
2. **Gallery Generation**: Executes the Go indexer to process photos
3. **Commit Changes**: Auto-commits generated gallery files
4. **Deploy**: Publishes the complete site to GitHub Pages

### File Structure

```
├── .github/workflows/
│   └── deploy.yml    # CI/CD workflow
├── web/                       # Website root
│   ├── photos/                 # Gallery photos
│   │   └── description.json    # Photo descriptions
│   ├── no-index-photos/        # Non-gallery photos (hero + featured)
│   ├── index/                  # Generated gallery JSON files
│   ├── css/                    # Stylesheets
│   ├── js/                     # JavaScript
│   └── index.html             # Main page
├── scripts/                    # Build tools
│   ├── indexer.go             # Photo gallery generator
│   └── go.mod                 # Go module file
├── generate_description_gemini.ps1 # AI description generator (Windows/PowerShell)
├── generate_description_gemini.sh  # AI description generator (Linux/macOS)
└── README.md                  # This file
```

## 🚀 Getting Started

To create your own memorial website:

1. **Fork this repository** to your GitHub account
2. **Edit the content** in `web/index.html` to reflect your loved one
3. **Add photos** to the `web/photos/` directory
4. **Update descriptions** in `web/photos/description.json` manually, or use the AI script (see below).
5. **Push to main branch** - GitHub Actions will automatically build and deploy

### Adding Photos

1. Place image files in `web/photos/`
2. Add descriptions to `web/photos/description.json`:
```json
{
    "photo1.jpg": "Description of photo 1",
    "photo2.jpg": "Description of photo 2"
}
```
3. Push changes - the gallery will automatically update

## 🖼️ Generating Descriptions with AI

This project includes scripts to automatically generate photo descriptions using AI. Choose the script for your operating system:
- **Windows**: `generate_description_gemini.ps1` (PowerShell)
- **Linux/macOS**: `generate_description_gemini.sh` (Bash)

### Requirements

- The [Google Gemini CLI](https://github.com/google/gemini-cli) - Make sure it's installed and authenticated
- **Linux/macOS only**: `jq` for JSON processing
  - CachyOS/Arch: `sudo pacman -S jq`
  - Ubuntu/Debian: `sudo apt install jq`
  - macOS: `brew install jq`

### How to Use

#### Windows (PowerShell)

1. Place your new photos in the `web/photos/` directory
2. Open a PowerShell terminal at the root of the project
3. Run the script:
   ```powershell
   .\generate_description_gemini.ps1
   ```

#### Linux/macOS (Bash)

1. Place your new photos in the `web/photos/` directory
2. Open a terminal at the root of the project
3. Make sure the script is executable (first time only):
   ```bash
   chmod +x generate_description_gemini.sh
   ```
4. Run the script:
   ```bash
   ./generate_description_gemini.sh
   ```

### How It Works

The script will:
1. Find all photos in `web/photos/` that don't have descriptions yet
2. Call the Gemini API (vision) to generate heartwarming descriptions for each photo
3. Incrementally save descriptions to `web/photos/description.json` after each success
4. Retry failed photos with rate limit handling
5. Save any permanently failed photos to `failed_photos.txt` for manual review

You can customize the prompt inside the script by editing the `PROMPT_TEMPLATE` variable (PowerShell) or `PROMPT_TEMPLATE` constant (Bash).

## 📝 License & Reuse

**Anyone is free to reuse this code for their own memorial websites.**

This project is shared with love in hopes that it can help others create beautiful tributes for their beloved companions. Whether it's for a pet, family member, or friend, feel free to adapt this codebase for your own memorial needs.

## 🤝 Contributing & Feedback

This project, including this README, was created with the assistance of AI. We welcome contributions and feedback!

-   **Showcase Your Memorial**: If you use this template, we would be delighted to see it. Please [create a GitHub issue](https://github.com/mackcoding/renzo-memorial/issues/new) to share a link to your project!
-   **Contribute Changes**: Fixes and improvements are always welcome! Please feel free to submit a Pull Request with a detailed description of your changes.

## 🛠️ Technical Requirements

- **Go** (for running the photo indexer)
- **GitHub Pages** enabled on your repository
- **GitHub Actions** enabled for automated builds

## 💝 In Memory

Renzo - Forever in our hearts and memories. A life well-lived, a love well-shared.

---

*"Until one has loved an animal, a part of one's soul remains unawakened." - Anatole France*
