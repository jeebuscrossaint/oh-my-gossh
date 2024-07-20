
// i daily drive void linux i kinda need to implement this lol, though xbps-src is NOT supported by tapeworm

use std::process::Command;

pub fn det_void() -> Option<String> {
    let xbps_path = "/usr/bin/xbps-install";
    let xbps_query_path = "/usr/bin/xbps-query";
    let xbps_remove_path = "/usr/bin/xbps-remove";

    if Command::new(xbps_path).arg("--version").output().is_ok() {
        println!("Detected xbps at: {}", xbps_path);
        Some(xbps_path.to_string())
    } else if Command::new(xbps_query_path).arg("--version").output().is_ok() {
        println!("Detected xbps-query at: {}", xbps_query_path);
        Some(xbps_query_path.to_string())
    } else if Command::new(xbps_remove_path).arg("--version").output().is_ok() {
        println!("Detected xbps-remove at: {}", xbps_remove_path);
        Some(xbps_remove_path.to_string())
    } else {
        None
    }
}