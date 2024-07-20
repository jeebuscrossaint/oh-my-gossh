
// arch linux backend lol pacman go crazy also i am NOT writing an aur helper so for that ur gonan need one of their backends lmao i aint doin allat

use std::process::Command;

pub fn det_arch() -> Option<String> {
    let pacman_path = "/usr/bin/pacman";
    let yay_path = "/usr/bin/yay";
    let paru_path = "/usr/bin/paru";

    if Command::new(paru_path).arg("--version").output().is_ok() {
        println!("Detected paru at: {}", paru_path);
        Some(paru_path.to_string())
    } else if Command::new(yay_path).arg("--version").output().is_ok() {
        println!("Detected yay at: {}", yay_path);
        Some(yay_path.to_string())
    } else if Command::new(pacman_path).arg("--version").output().is_ok() {
        println!("Detected pacman at: {}", pacman_path);
        Some(pacman_path.to_string())
    } else {
        None
    }
}