# renpho

renpho renames your photo files with information from its exif info

Finding photos is difficult when each picture is named DSC-xxx.
renpho will rename these files and adds the date and time from the photo exif data
as the filename AND the files timestamp. So all your photos will be sorted according to their creation timestamp.

JPG and RAW files are moved to different directories.

```bash
renpho -l "Xmas Party" -dryrun -p /demo

░░░░░░░░░░░░░░░░░
░ Rename Photos ░
░░░░░░░░░░░░░░░░░

❕ INFO: Path : /demo
❕ INFO: Label: Xmas-Party
❕ INFO: ======================================
🏁 YAY - Dryrun - won`t change anything
💣 EROR: in file CK21_st090.JPG - Defaulting to 'unknown' (exif: tag "Model" is not present)
💣 EROR: in file CK21_st090.JPG - Defaulting to 'TODAY' (exif: tag "DateTime" is not present)
❕ INFO: moving CK21_st090.JPG -> /2021-12-12_Xmas-Party_JPG/2021-12-12-13h00m10_unknown_CK21_st090.JPG
❕ INFO: moving DSCF0001.JPG -> /2021-12-12_Xmas-Party_JPG/2020-12-01-21h38m49_X-Pro3_DSCF0001.JPG
❕ INFO: moving DSCF3523.JPG -> /2021-12-12_Xmas-Party_JPG/2019-01-18-00h01m11_X-T1_DSCF3523.JPG
❕ INFO: moving LDN_9186.JPG -> /2021-12-12_Xmas-Party_JPG/2021-11-19-12h18m12_NIKON-D750_LDN_9186.JPG
🏁 YAY - All done :)
```

Don't hit me - it's my first golang project... ^\_^v
