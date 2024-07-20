// goofy proprietary system but whatevs ibm u do u

use std::path::Path;

pub fn det_redhat() -> Option<String> {
    let dnf_path = "/usr/bin/dnf";
    let rpm_path = "/usr/bin/rpm";
    let yum_path = "/usr/bin/yum";

    if Path::new(dnf_path).exists() {
        println!("Detected dnf at: {}", dnf_path);
        Some(dnf_path.to_string())
    } else if Path::new(rpm_path).exists() {
        println!("Detected rpm at: {}", rpm_path);
        Some(rpm_path.to_string())
    } else if Path::new(yum_path).exists() {
        println!("Detected yum at: {}", yum_path);
        Some(yum_path.to_string())
    } else {
        None
    }
}