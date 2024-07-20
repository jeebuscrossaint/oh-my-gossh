// never used fedora ever never plan too but it seems cool

use std::process::Command;

pub fn det_fedora() -> Option<String> {
    let dnf_path = "/usr/bin/dnf";
    let rpm_path = "/usr/bin/rpm";
    let yum_path = "/usr/bin/yum";

    if Command::new(dnf_path).arg("--version").output().is_ok() {
        println!("Detected dnf at: {}", dnf_path);
        Some(dnf_path.to_string())
    } else if Command::new(rpm_path).arg("--version").output().is_ok() {
        println!("Detected rpm at: {}", rpm_path);
        Some(rpm_path.to_string())
    } else if Command::new(yum_path).arg("--version").output().is_ok() {
        println!("Detected yum at: {}", yum_path);
        Some(yum_path.to_string())
    } else {
        None
    }
}