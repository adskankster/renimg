# renimg

Run the application to change the filenames of imagefiles to that of the Original Date tag. This enables the files to be sorted by date and time of taking.

$ renimg [--dir=<path>] [-d|--dry-run] [--debug]

--dir           the directory to work on. Default is current directory. Will not drill down to sub directories.

--dry-run -d    show the changes that would be made, but do not make them.

--debug         output debug messages.


The application will skip files that are already in the destination format.

----

Project to learn Go. I decided to replace a bash script I had the used Exiftool to convert image filenames form XXXXXXYYY.jpg to YYYY-MM-DD-HH-MM-SS-XXXXXXYYY.jpg (actual initial filename can vary, and at the moment the application will only deal with jpeg extensionss). The date and time used is based on the Original Date tag, but could be changed to use Create Date.

I used RWCarlsen's exif library, and used MRauh's Exifsorter as a starting point. My use case is slghtly different from Exifsorter and I needed the learning exercise.

I might update this application to provide more options, but the bash script was in use for several years without needing to be modified, so my current usage is probably satisfied.

Constructive advice on coding practice is welcome.

With thanks to goexif library: "github.com/rwcarlsen/goexif"

With thanks to Exifsorter for providing a template to start from.
See: github.com/mrauh/exifsorter

