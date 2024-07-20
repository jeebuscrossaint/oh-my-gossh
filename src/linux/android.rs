
// i use android and termux on a daily basis my screentime is wild its prob my second most used app for some reason idk lol

use std::process::Command;

pub fn det_termux() -> Option<String> {
    let pkg_path = "/data/data/com.termux/files/usr/bin/pkg";
    let dpkg_path = "/data/data/com.termux/files/usr/bin/dpkg";
    let apt_path = "/data/data/com.termux/files/usr/bin/apt";

    if Command::new(pkg_path).arg("list").output().is_ok() {
        println!("Detected pkg at: {}", pkg_path);
        Some(pkg_path.to_string())
    } else if Command::new(dpkg_path).arg("--version").output().is_ok() {
        println!("Detected dpkg at: {}", dpkg_path);
        Some(dpkg_path.to_string())
    } else if Command::new(apt_path).arg("--version").output().is_ok() {
        println!("Detected apt at: {}", apt_path);
        Some(apt_path.to_string())
    } else {
        None
    }
}