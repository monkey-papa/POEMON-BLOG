# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.sort import Sort

class SaveSortView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            sortName = data['sortName']
            sortDescription = data['sortDescription']
            sortType = data['sortType']
            priority = data['priority']
            if sortName and sortDescription and priority and sortType in [0, 1]:
                Sort.objects.create(sort_name=sortName, sort_description=sortDescription,
                                    sort_type=sortType, priority=priority)
                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null or error",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
