import sys
from PIL import Image
import exifread

def extract_gps_info(image_path):
    with open(image_path, 'rb') as f:
        tags = exifread.process_file(f)
        gps_lat = tags.get('GPS GPSLatitude')
        gps_lat_ref = tags.get('GPS GPSLatitudeRef')
        gps_lon = tags.get('GPS GPSLongitude')
        gps_lon_ref = tags.get('GPS GPSLongitudeRef')

        if gps_lat and gps_lon and gps_lat_ref and gps_lon_ref:
            lat = convert_to_degrees(gps_lat)
            lon = convert_to_degrees(gps_lon)
            if gps_lat_ref.values[0] != 'N':
                lat = -lat
            if gps_lon_ref.values[0] != 'E':
                lon = -lon
            return lat, lon
        else:
            return None, None

def convert_to_degrees(value):
    d, m, s = value.values
    return round(d.num / d.den + (m.num / m.den) / 60 + (s.num / s.den) / 3600, 4)

def extract_hidden_data(image_path):
    image = Image.open(image_path)
    pixels = list(image.getdata())
    binary_string = ''
    for pixel in pixels:
        for color in pixel[:3]:  # Only the RGB values
            binary_string += bin(color)[-1]  # Take the LSB of each color value

    byte_size = 8
    hidden_data = ''
    for i in range(0, len(binary_string), byte_size):
        byte = binary_string[i:i + byte_size]
        hidden_data += chr(int(byte, 2))
    return hidden_data

def main(image_path, option):
    if option == '-map':
        lat, lon = extract_gps_info(image_path)
        if lat and lon:
            print(f'Lat/Lon:\t({lat}) / ({lon})')
        else:
            print('No GPS data found')
    elif option == '-steg':
        hidden_data = extract_hidden_data(image_path)
        print("-----BEGIN PGP PUBLIC KEY BLOCK-----" + "\n" + hidden_data.strip()
               + "\n" + "-----END PGP PUBLIC KEY BLOCK-----")
    else:
        print('Invalid option. Use -map for GPS data or -steg for hidden data.')

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print('Usage: python inspector_image.py [option] [image_path]')
        sys.exit(1)
    
    option = sys.argv[1]
    image_path = sys.argv[2]
    main(image_path, option)
