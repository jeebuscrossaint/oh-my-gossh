mod detect_sys;
mod linux_spec;
mod bsd_flav;

fn main() {
    let os_name = detect_sys::detect_class::find_os();
    if os_name == "Linux" {
        let distro = detect_sys::detect_distro::find_distro().expect("Failed to find distro");
        if distro == "debian" {
            linux_spec::debian::det_debian();
        } else if distro == "arch" {
            linux_spec::arch::det_arch();
        } else if distro == "alpine" {
            linux_spec::alpine::det_alpine();
        } else if distro == "android" {
            linux_spec::android::det_termux();
        } else if distro == "fedora" {
            linux_spec::fedora::det_fedora();
        } else if distro == "gentoo" {
            linux_spec::gentoo::det_gentoo();
        } else if distro == "redhat" {
            linux_spec::redhat::det_redhat();
        } else if distro == "suse" {
            linux_spec::suse::det_suse();
        } else if distro == "slackware" {
            linux_spec::slackware::det_slackware();
        } else if distro == "void" {
            linux_spec::void::det_void();
        } else {
            println!("Unsupported Distro");
        }
    } else if os_name == "Windows" {
        detect_sys::detect_win::det_win();
    } else if os_name == "macOS" {
        detect_sys::detect_mac::det_mac();
    } else if os_name == "BSD" {
        detect_sys::detect_flavor::detect_bsd_variant().expect("Failed to detect BSD variant");
        if os_name == "FreeBSD" {
            bsd_flav::free::det_freebsd();
        } else if os_name == "OpenBSD" {
            bsd_flav::open::det_openbsd();
        } else if os_name == "NetBSD" {
            bsd_flav::net::det_netbsd();
        } else if os_name == "DragonFly" {
            bsd_flav::dragonfly::det_dragonfly();
        } else {
            println!("Unsupported BSD variant");
        }
    } else if os_name == "Android" {
        detect_sys::detect_termux::det_termux();
        linux_spec::android::det_termux();
    }
    else {
        println!("Unsupported OS");
    }
}