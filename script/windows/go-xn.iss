﻿; Script generated by Huiyi.FYJ
; 2019 | xn-02f Lab

; ID used to uniquely identify this application.
; Generated by PowerShell:
;     [guid]::NewGuid().ToString()
#define APP_ID '9F00A778-1C16-4F7D-8FE4-0CDE4FC712DD'
#define APP_NAME 'Go-xn'
#define VERSION '0.0.2'
#define PUBLISHER 'xn-02f Lab'
#define URL 'https://xn--02f.com'
#define EXE_NAME 'go-xn.exe'
#define ICON '..\..\web\icons\xn-02f.ico'
#define LICENSE '..\..\LICENSE'
#define OUTPUT_DIR '..\..\build\dist'
#define OUTPUT_NAME 'go-xn-setup'

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={#APP_ID}
AppName={#APP_NAME}
AppVersion={#VERSION}
AppVerName={#APP_NAME} {#VERSION}
AppPublisher={#PUBLISHER}
AppPublisherURL={#URL}
AppSupportURL={#URL}
AppUpdatesURL={#URL}
DefaultDirName=C:\{#APP_NAME}
LicenseFile={#LICENSE}
OutputDir={#OUTPUT_DIR}
OutputBaseFilename={#OUTPUT_NAME}
; better size is 164x314
WizardImageFile=..\assets\inno-big.bmp
; better size is 55x55
WizardSmallImageFile=..\assets\inno-small.bmp
SetupIconFile={#ICON}
UninstallDisplayIcon={#ICON}
Uninstallable=yes
; Remove the following line to run in admin mode (install for all users).
; 'PrivilegesRequiredOverridesAllowed=dialog' can alter mode by dialog.
PrivilegesRequired=lowest
DisableProgramGroupPage=yes
DisableReadyPage=no
DisableDirPage=no
DirExistsWarning=yes
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl,.\i18n\Messages.en.isl"
Name: "chinesesimplified"; MessagesFile: ".\i18n\ChineseSimplified.isl,.\i18n\Messages.zh-cn.isl"

[Files]
Source: "..\..\build\go-xn.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\..\web\*"; DestDir: "{app}\web"; Flags: ignoreversion recursesubdirs createallsubdirs
Source: "..\..\LICENSE"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\..\README.md"; DestDir: "{app}"; Flags: ignoreversion isreadme
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:Other}"; Flags: unchecked
Name: "addtopath"; Description: "{cm:AddToPath}"; GroupDescription: "{cm:Other}"; Flags: unchecked

[Icons]
Name: "{autoprograms}\{#APP_NAME}"; Filename: "{app}\{#EXE_NAME}"
Name: "{autodesktop}\{#APP_NAME}"; Filename: "{app}\{#EXE_NAME}"; Comment: "{cm:DesktopIconComment}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#EXE_NAME}"; Parameters: "web"; Description: "{cm:LaunchProgram,{#StringChange(APP_NAME, '&', '&&')}}"; Flags: nowait postinstall skipifsilent
