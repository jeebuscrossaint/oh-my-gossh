// never used suse before but ok

use std::process::Command;

pub fn det_suse() -> Option<String> {
    let zypper_path = "/usr/bin/zypper";
    let rpm_path = "/usr/bin/rpm";

    if Command::new(zypper_path).arg("--version").output().is_ok() {
        println!("Detected zypper at: {}", zypper_path);
        Some(zypper_path.to_string())
    } else if Command::new(rpm_path).arg("--version").output().is_ok() {
        println!("Detected rpm at: {}", rpm_path);
        Some(rpm_path.to_string())
    } else {
        None
    }
}