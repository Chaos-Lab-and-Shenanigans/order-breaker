# 🌟 Astrology Rickroller - The Cosmic Deceiver

> **"You expect ancient wisdom, but you get something better: Eternal Rickroll."**

![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go)
![UI-Toolkit](https://img.shields.io/badge/UI%20Toolkit-Fyne-296E9E?style=for-the-badge&logo=go)
![Status](https://img.shields.io/badge/Prank%20Status-Operational-green?style=for-the-badge)
![Latest Release](https://img.shields.io/github/v/release/Chaos-Lab-and-Shenanigans/astrology-rickroller?color=orange&label=Latest%20Release&style=for-the-badge)

## 🌌 About The Project

The **Astrology Rickroller** is a cross-platform desktop application designed to deceive and delight. It is expertly disguised as a serious utility for calculating and displaying personalized horoscopes based on the user's Zodiac sign.

### The Shenanigans

1.  The user opens the beautifully crafted Fyne application.
2.  They are prompted to enter information and click a button to receive their reading.
3.  Upon submission, the app displays a convincing "Calculating Star Charts..." message.
4.  Then roasts the user and displays a tempting "See interesting information" button. 
5.  Instead of a horoscope, the application plays the unforgettable audio track of **"Never Gonna Give You Up"** at maximum volume, renaming the desktop files to lyrics, changing the wallpaper,  rickrolling the user through their speakers.
6.  The user is stuck in the app after getting rickrolled, you think exit button will work? Try and see it for yourself. If user closes the app, the app reappears every 5 seconds, absolute cinema.

## Recover from chaos

- Go to home page.  
- Go to compatability checker.
- Set your own DOB manually to **`01/01/6969`**(The other one doesn't matter).
- Cick on "See interesting information" button to recover.  

## ✨ Features

* **Native Desktop Experience:** Built with the Fyne UI toolkit for a clean, professional look.
* **Audio And Wallpaper Payload:** Uses the `assets` folder to deliver a high-quality audio and wallpaper rickroll, which are embedded in app.
* **Easy Distribution:** Easily built into platform-specific executables.

## 🚀 Getting Started

There are two ways to get the Astrology Rickroller running:

### 1. For the End User

Download a pre-built executable from the **Releases** page:

1.  Go to the [Releases Tab](https://github.com/Chaos-Lab-and-Shenanigans/astrology-rickroller/releases) of this repository.
2.  Download the latest version.
3.  Run the downloaded executable and share it with your unsuspecting friends!

### 2. For the Go Developer (Building from Source)

If you have the Go toolchain installed, you can build the application yourself:

1.  **Clone the Repository:**
    ```bash
    git clone [https://github.com/Chaos-Lab-and-Shenanigans/astrology-rickroller.git](https://github.com/Chaos-Lab-and-Shenanigans/astrology-rickroller.git)
    cd astrology-rickroller
    ```

2.  **Install Fyne Dependencies:**
    * *Note: Fyne requires standard system dependencies. Check the [Fyne Documentation](https://developer.fyne.io/started/) if you encounter build errors.*

3.  **Build and Run:**
   
    Make sure you have `fyne-cross` installed by running:     
    ```bash
    go install github.com/fyne-io/fyne-cross@latest
    ```

    Use fyne-cross to build and run. For example for windows:  
    ```bash
    fyne-cross windows
    ```

## 🛠 Customization: Tailoring the Chaos

The project is structured to allow for easy changes to the look, sound, and core deception parameters without touching the main application logic (`main.go`).

### 1. Visual Customization (Assets)

| Component | File Path | Action |
| :--- | :--- | :--- |
| **Wallpaper/Background** | `assets/wall.png` | Replace this file to change the background image that gets set as a wallpaper. |
| **Audio Prank Payload** | `assets/audio.mp3` | Replace this file to use a different audio clip or song for the rickroll. |
| **Icon** | `Icon.png` | Replace this file to change the application icon. |

### 2. Deep Configuration (The Code)

The `internal/config/config.go` file acts as the central hub for nearly all application text and parameter settings.

| Setting | File Path | Detail |
| :--- | :--- | :--- |
| **The Lyrics/Text** | `internal/config/config.go` | Change the string array variables in this file to modify the actual text/lyrics that rename the desktop files. |
| **Window Size & App Identity** | `internal/config/config.go` | Modify settings like the default window dimensions, the application name (`astrology.exe`), and other UI-related variables. |
| **"Chaos" Parameters** | `internal/config/config.go` | Adjust values like `DateForRestore` used to "recover the chaos". |

***
## 🤝 Contributing

We welcome contributions that improve the deception, stability, and cross-platform compatibility!

1.  Fork the Project.
2.  Create your Feature Branch (`git checkout -b feature/BetterZodiacUI`).
3.  Commit your Changes (`git commit -m 'feat: improved Fyne layout for better deception'`).
4.  Push to the Branch (`git push origin feature/BetterZodiacUI`).
5.  Open a Pull Request.

---

<p align="center">
  A highly focused project built with 💖 and 😈 by <a href="https://github.com/Chaos-Lab-and-Shenanigans">Chaos Lab and Shenanigans</a>
</p>
