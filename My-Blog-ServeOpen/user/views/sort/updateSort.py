# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.sort import Sort


class UpadateSortView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            sortName = data['sortName']
            sortDescription = data['sortDescription']
            sortType = data['sortType']
            priority = data['priority']
            if id and sortName and sortDescription and priority and sortType in [0, 1]:
                Sort.objects.filter(id=id).update(sort_name=sortName, sort_description=sortDescription,
                                                  sort_type=sortType, priority=priority)
                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
