#!/usr/bin/env python
from collections import namedtuple
from ansible.parsing.dataloader import DataLoader
from ansible.vars import VariableManager
from ansible.inventory import Inventory
from ansible.playbook.play import Play
from ansible.executor.task_queue_manager import TaskQueueManager

from ansible.plugins.callback import CallbackBase
import json
from rest_framework import viewsets
from rest_framework.decorators import detail_route, list_route
from rest_framework.response import Response


class ResultCallback(CallbackBase):
    """A sample callback plugin used for performing an action as results come in

    If you want to collect all results into a single object for processing at
    the end of the execution, look into utilizing the ``json`` callback plugin
    or writing your own custom callback plugin
    """

    def v2_runner_on_ok(self, result, **kwargs):
        """Print a json representation of the result

        This method could store the result in an instance attribute for retrieval later
        """
        host = result._host
        self.datas = json.dumps({host.name: result._result}, indent=4)


results_callback = ResultCallback()

host = '192.168.0.104'
Options = namedtuple('Options', ['listtags', 'listtasks', 'listhosts', 'syntax', 'connection', 'module_path', 'forks',
                                 'private_key_file', 'ssh_common_args', 'ssh_extra_args',
                                 'sftp_extra_args', 'scp_extra_args', 'become', 'become_method', 'become_user',
                                 'verbosity', 'check'])
variable_manager = VariableManager()
loader = DataLoader()

options = Options(listtags=False, listtasks=False, listhosts=False, syntax=False, connection='ssh', module_path=None,
                  forks=100, private_key_file=None,
                  ssh_common_args=None, ssh_extra_args=None, sftp_extra_args=None, scp_extra_args=None, become=False,
                  become_method=None, become_user=None, verbosity=None, check=False)
passwords = {}

inventory = Inventory(loader=loader, variable_manager=variable_manager)
variable_manager.set_inventory(inventory)
play_source = dict(
    name="Ansible Play",
    hosts="192.168.0.104",
    gather_facts='no',
    tasks=[
        dict(action=dict(module='shell', args='ls'), register='shell_out')
    ]
)
play = Play().load(play_source, variable_manager=variable_manager, loader=loader)
qm = None

#@list_route(methods=['get', 'post'])
'''
def run_ansible(self,request):
  try:
    tqm = TaskQueueManager(
        inventory=inventory,
        variable_manager=variable_manager,
        loader=loader,
        options=options,
        passwords=passwords,
        stdout_callback=results_callback,
    )
    result = tqm.run(play)

    jason_result = results_callback.datas
    return Response(jason_result)


    # exit()

# with open("ansible_exec_result.html", 'rw') as files:
#       files.write(data)
#        string = files.read(data)


  finally:
    if tqm is not None:
        tqm.cleanup()
'''
