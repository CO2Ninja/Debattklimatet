from django.views.generic import ListView
from models import twitteruser
from django.db.models import Max

class TwitteruserDetail(ListView):

    model = twitteruser

    def get_context_data(self, **kwargs):
        # Call the base implementation first to get a context
        context = super(ListView, self).get_context_data(**kwargs)
        # Add in a QuerySet of all the books
        context['maxscore'] = twitteruser.objects.aggregate(Max('totalscore'))['totalscore__max']
        return context
