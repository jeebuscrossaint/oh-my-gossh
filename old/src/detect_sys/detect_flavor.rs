/// couldnt leave out my bsd enjoyers roots

use std::process::Command;
use std::str;

pub fn detect_bsd_variant() -> Option<String> {
    let output = Command::new("uname")
        .arg("-s")
        .output()
        .expect("failed to execute process");

    let sysname = str::from_utf8(&output.stdout).ok()?.trim().to_string();

    match sysname.as_str() {
        "FreeBSD" => Some("FreeBSD".to_string()),
        "OpenBSD" => Some("OpenBSD".to_string()),
        "NetBSD" => Some("NetBSD".to_string()),
        "DragonFly" => Some("DragonflyBSD".to_string()),
        _ => None,
    }
}