<#
    .SYNOPSIS
    Installs zagent-vm

    .DESCRIPTION
    Installs zagent-vm to $HOME\zagent
    If zagent-vm is already installed, try to update to the latest version.

    .PARAMETER Uninstall
    Uninstall zagent-vm. Note that this uninstalls any Python versions that were installed with zagent-vm.

    .INPUTS
    None.

    .OUTPUTS
    None.

    .EXAMPLE
    PS> install-zagent-vm.ps1

    .LINK
    Online version: https://github.com/easysoft/zenagent
#>
    
param (
    $s,$z
)

$AgentDir = "${env:USERPROFILE}\zagent"
$AgentWinDir = "${AgentDir}\zagent-win"

Function Get-CurrentVersion() {
    $VersionFilePath = "$AgentDir\latest.version"
    If (Test-Path $VersionFilePath) {
        $CurrentVersion = Get-Content $VersionFilePath
    }
    Else {
        $CurrentVersion = ""
    }

    Return $CurrentVersion
}

Function Get-LatestVersion() {
    $LatestVersionFilePath = "$AgentDir\latest.version"
    If (-not (Test-Path $AgentDir)) {
        New-Item -Path $AgentDir -ItemType Directory
    }
    (New-Object System.Net.WebClient).DownloadFile("https://pkg.qucheng.com/zenagent/app/windows/vm.version", $LatestVersionFilePath)
    $LatestVersion = Get-Content $LatestVersionFilePath

    Return $LatestVersion
}

Function Main() {
    $CurrentVersion = Get-CurrentVersion
    If ($CurrentVersion) {
        Write-Host "zagent-vm $CurrentVersion installed."
    }
    $LatestVersion = Get-LatestVersion
    If ($CurrentVersion -eq $LatestVersion) {
        Write-Host "No updates available."
        if((Get-Process -Name vm -ErrorAction SilentlyContinue) -ne $null) {
            Write-Host "Kill running zagent-vm.exe"
            Stop-Process -Name vm
        }
        Write-Host "starting server by $AgentDir\zagent-vm.exe -p 55201 -secret $s -s $z"
        Start-Process -WorkingDirectory $AgentDir -FilePath "zagent-vm.exe" -ArgumentList "-p 55201 -secret $s -s $z"
        exit
    }
    Else {
        If ($CurrentVersion) {
            Write-Host "New version available: $LatestVersion. Updating..."
        }
    }   

    
    If (-not (Test-Path $AgentDir)) {
        New-Item -Path $AgentDir -ItemType Directory
    }

    $DownloadPath = "$AgentDir\zagent-vm.zip"

    Write-Host "Downloading file..."
    (New-Object System.Net.WebClient).DownloadFile("https://pkg.qucheng.com/zenagent/app/windows/vm.zip", $DownloadPath)
    Write-Host "Unzip file..."
    Microsoft.PowerShell.Archive\Expand-Archive -Path $DownloadPath -DestinationPath $AgentDir
    Move-Item -Path "$AgentDir\vm\vm.exe" -Destination "$AgentDir\zagent-vm.exe"
    Remove-Item -Path "$AgentDir\vm" -Recurse
    Remove-Item -Path $DownloadPath


    $StartDir = "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup\"

    If (-not (Test-Path "$StartDir\zagent-vm.bat")) {
        Write-Host "Add zagent-vm.bat to $StartDir"
        "@echo on`ncmd /k `"cd /d $AgentDir && zagent-vm.exe -p 55201 -secret $s -s $z`"`npause" | Out-File "$AgentDir\zagent-vm.bat" -Encoding ASCII
        If(-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator"))
        {   
            "Move-Item -Path `"$AgentDir\zagent-vm.bat`" -Destination `"$StartDir`"" | Out-File "$AgentDir\zagent-vm-start.ps1"
            Start-Process "$psHome\powershell.exe" -ArgumentList "$AgentDir\zagent-vm-start.ps1" -verb runas
        }else{
            Move-Item -Path "$AgentDir\zagent-vm.bat" -Destination "$StartDir"
        }
    }

    Write-Host "starting server by $AgentDir\zagent-vm.exe -p 55201 -secret $s -s $z"
    Start-Process -WorkingDirectory $AgentDir -FilePath "zagent-vm.exe" -ArgumentList "-p 55201 -secret $s -s $z"

    If ($? -eq $True) {
        Write-Host "zagent-vm is successfully installed"
    }
    Else {
        Write-Host "zagent-vm was not installed successfully"
    }
}

Main