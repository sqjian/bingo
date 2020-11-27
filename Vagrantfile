Vagrant.configure('2') do |config|
  config.vm.box = "ubuntu/focal"
  config.vm.network :private_network, type: "dhcp"
  config.vm.network "forwarded_port", guest: 80, host: 80
  config.vm.box_check_update = false
  config.vm.hostname = "testing"

  config.vm.provider :virtualbox do |vb, override|
      vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      vb.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
      vb.customize ["modifyvm", :id, "--memory", "8192"]
      vb.customize ["modifyvm", :id, "--cpus", 4]
  end
  
  config.vm.provision "shell", name: "docker", path: "scripts/docker.sh"
  config.vm.provision "shell", name: "alias", path: "scripts/alias.sh"

  config.ssh.username = "root"
  config.ssh.password = "9527"
end