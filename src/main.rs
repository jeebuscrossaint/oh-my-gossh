mod detect_sys;

fn main() {
    let os_name = detect_sys::detect_class::find_os();
    if os_name == "Linux" {
        detect_sys::detect_distro::find_distro().expect("Failed to find distro");
    } else if os_name == "Windows" {
        detect_sys::detect_win::det_win();
    } else if os_name == "macOS" {
        detect_sys::detect_mac::det_mac();
    } else if os_name == "BSD" {
        detect_sys::detect_flavor::detect_bsd_variant();
    } else if os_name == "Android" {
        detect_sys::detect_termux::det_termux();
    }
     else {
    println!("test");
    }
}