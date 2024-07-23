
// net bsd backend pkgsrc

use std::process::Command;

pub fn det_netbsd() -> Option<String> {
    let pkgin_path = "/usr/sbin/pkgin";
    let pkgsrc_path = "/usr/sbin/pkgsrc";

    if Command::new(pkgin_path).arg("--version").output().is_ok() {
        println!("Detected pkgin at: {}", pkgin_path);
        Some(pkgin_path.to_string())
    } else if Command::new(pkgsrc_path).arg("--version").output().is_ok() {
        println!("Detected pkgsrc at: {}", pkgsrc_path);
        Some(pkgsrc_path.to_string())
    } else {
        None
    }
}