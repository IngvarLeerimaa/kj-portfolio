# [Inspector Image](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/inspector-image)

## Objective
- Create a program that will inspect an image and return the image's metadata.

This lesson mainly focuses on steganography and hidden metadata.

Steganography is the practice of concealing messages, images, or files within other, seemingly ordinary, messages, images, or files. Unlike cryptography, which obscures the content, steganography hides its very existence.

Data can be hidden in normal files using various steganography techniques:
- **Image Files**: Embed data in the least significant bits (LSBs) of pixel values, so the change is imperceptible.
- **Audio Files**: Modify the LSBs of audio samples or use phase coding and spread spectrum techniques.
- **Text Files**: Alter formatting, such as spacing or font size, or use character encoding methods.
- **Video Files**: Combine methods used for images and audio in individual frames.
- **Network Protocols**: Embed data in unused fields or manipulate packet timing.

Each method ensures the hidden data is difficult to detect without altering the apparent normal file significantly.

### Setup
1. Install the requirements using:

    ```bash
    pip install -r requirements.txt
    ```

### Usage

1. To get the latitude and longitude of the image, run the following command: 

    ```bash
    python3 inspector_image.py -map image.jpeg
    ```

    For this image.jpeg, the output will be:

    ```python
    Lat/Lon: (32.0866) / (34.8851)
    ```

2. To get the metadata of the image, run the following command:

    ```bash
    python3 inspector_image.py -steg image.jpeg
    ```

    For this image.jpeg, the output will be:

    ```plaintext
    -----BEGIN PGP PUBLIC KEY BLOCK-----
    [pgp key]
    -----END PGP PUBLIC KEY BLOCK-----
    ```

### Technologies used
- Pillow
- Exifread

### [Task Description](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/inspector-image)
### [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/inspector-image/audit)
