
// portage is pretty fire im fond of gentoo despite never daily driving it

use std::process::Command;

pub fn det_gentoo() -> Option<String> {
    let emerge_path = "/usr/bin/emerge";
    let eix_path = "/usr/bin/eix";

    if Command::new(emerge_path).arg("--version").output().is_ok() {
        println!("Detected emerge at: {}", emerge_path);
        Some(emerge_path.to_string())
    } else if Command::new(eix_path).arg("--version").output().is_ok() {
        println!("Detected eix at: {}", eix_path);
        Some(eix_path.to_string())
    } else {
        None
    }
}