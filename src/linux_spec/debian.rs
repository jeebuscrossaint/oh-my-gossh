
// debian backend basically aptitude/dpkg backend

use std::process::Command;

pub fn det_debian() -> Option<String> {
    let apt_path = "/usr/bin/apt";
    let dpkg_path = "/usr/bin/dpkg";

    if Command::new(apt_path).arg("--version").output().is_ok() {
        println!("Detected apt at: {}", apt_path);
        Some(apt_path.to_string())
    } else if Command::new(dpkg_path).arg("--version").output().is_ok() {
        println!("Detected dpkg at: {}", dpkg_path);
        Some(dpkg_path.to_string())
    } else {
        None
    }
}

