use std::path::Path;
use std::process::Command;

pub fn det_win() -> String {
    let pm_name = if is_scoop_available() {
        "Scoop"
    } else if is_choco_available() {
        "Chocolatey"
    } else if is_winget_available() {
        "Winget"
    } else {
        "No known package manager found"
    };

    println!("Detected package manager: {}", pm_name);
    pm_name.to_string()
}

fn is_scoop_available() -> bool {
    let scoop_path = dirs::home_dir().map(|p| p.join("scoop")).unwrap_or_default();
    scoop_path.exists() || Command::new("scoop").arg("list").output().is_ok()
}

fn is_choco_available() -> bool {
    let choco_path = Path::new("C:\\ProgramData\\chocolatey");
    choco_path.exists() || Command::new("choco").arg("-v").output().is_ok()
}

fn is_winget_available() -> bool {
    Command::new("winget").arg("--version").output().is_ok()
}