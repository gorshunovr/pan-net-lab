Vagrant.configure("2") do |config|
    # Ubuntu 16.04 Xenial x64
    # Box supports most of the providers, including libvirt and virtualbox
    config.vm.box = "generic/ubuntu1604"
    config.vm.provider :libvirt do |v|
      v.qemu_use_session = false # BZ1697773
                                 # "bug" with private networks not being created
    end
#    config.vm.provider "virtualbox"
#    config.vm.provider "libvirt"

    # main node
    config.vm.define "n0" do |n0|
      n0.vm.hostname = "n0"
      config.vm.provider "libvirt" do |v, override|      # rsync for libvirt
        override.vm.synced_folder './', '/vagrant', type: 'rsync'
      end
      n0.vm.network :private_network, :ip => "192.168.56.10"
      n0.vm.provision :shell, path: "bootstrap.sh" # ansible installation
      n0.vm.provision :shell, inline: "cd /vagrant/ansible/ && ansible-playbook -i inventory.txt setup.yml"
    end

    # test node
    config.vm.define "n1" do |n1|
      n1.vm.hostname = "n1"
      n1.vm.network :private_network, :ip => "192.168.56.11"
      [:virtualbox, :parallels, :libvirt, :hyperv].each do |provider|
        n1.vm.provider provider do |vplh, override|
          vplh.cpus = 1
          vplh.memory = 512
        end
      end
      [:vmware_fusion, :vmware_workstation, :vmware_desktop].each do |provider|
        n1.vm.provider provider do |vmw, override|
          vmw.vmx["numvcpus"] = "1"
          vmw.vmx["memsize"] = "512"
        end
      end
    end
end
