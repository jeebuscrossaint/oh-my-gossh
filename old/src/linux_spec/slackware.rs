// oldhead distro but alright it exists gotta acknowledge it and its an original distro too

use std::path::Path;

pub fn det_slackware() -> Option<String> {
    let slackpkg_path = "/usr/sbin/slackpkg";
    let pkgtool_path = "/sbin/pkgtool";

    if Path::new(slackpkg_path).exists() {
        println!("Detected slackpkg at: {}", slackpkg_path);
        Some(slackpkg_path.to_string())
    } else if Path::new(pkgtool_path).exists() {
        println!("Detected pkgtool at: {}", pkgtool_path);
        Some(pkgtool_path.to_string())
    } else {
        None
    }
}