[Setup]
AppName=pypi-versions
AppVersion=1.1.13
DefaultDirName={pf}\pypi-versions
DefaultGroupName=pypi-versions
OutputDir={src}
OutputBaseFilename=pypi-versions-setup
Compression=lzma
SolidCompression=yes

[Files]
Source: "ppv.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\pypi-versions"; Filename: "{app}\ppv.exe"