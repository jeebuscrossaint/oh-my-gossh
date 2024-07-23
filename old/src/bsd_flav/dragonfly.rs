
// dragonfly bsd backend dports

use std::process::Command;

pub fn det_dragonfly() -> Option<String> {
    let dports_path = "/usr/sbin/dports";
    let pkg_path = "/usr/sbin/pkg";

    if Command::new(dports_path).arg("--version").output().is_ok() {
        println!("Detected dports at: {}", dports_path);
        Some(dports_path.to_string())
    } else if Command::new(pkg_path).arg("--version").output().is_ok() {
        println!("Detected pkg at: {}", pkg_path);
        Some(pkg_path.to_string())
    } else {
        None
    }
}