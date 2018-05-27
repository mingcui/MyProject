from django.shortcuts import render
from django.http import HttpResponse
from rest_framework import viewsets
from rest_framework.response import Response
from .ansible_class import *

from rest_framework.decorators import detail_route, list_route


# Create your views here.
#class Run_ansible(viewsets.ModelViewSet):
  #  @list_route(methods=['get', 'post'])
def run_df(request):
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
            return HttpResponse(jason_result)
       finally:
            if tqm is not None:
                tqm.cleanup()
