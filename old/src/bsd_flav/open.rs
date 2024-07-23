
// openbsd backend

use std::process::Command;

pub fn det_openbsd() -> Option<String> {
    let pkg_path = "/usr/sbin/pkg";
    let pkgng_path = "/usr/sbin/pkgng";

    if Command::new(pkg_path).arg("--version").output().is_ok() {
        println!("Detected pkg at: {}", pkg_path);
        Some(pkg_path.to_string())
    } else if Command::new(pkgng_path).arg("--version").output().is_ok() {
        println!("Detected pkgng at: {}", pkgng_path);
        Some(pkgng_path.to_string())
    } else {
        None
    }
}