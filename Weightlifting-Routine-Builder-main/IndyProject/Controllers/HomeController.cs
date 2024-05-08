using IndyProject.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;

namespace IndyProject.Controllers
{
    public class HomeController : Controller
    {
        [HttpGet]
        public IActionResult Index()
        {
            ViewBag.bmr = 0;
            ViewBag.orm = 0;
            ViewBag.lIntensityLoss = 0;
            ViewBag.mIntensityLoss = 0;
            ViewBag.hIntensityLoss = 0;
            ViewBag.lIntensityGain = 0;
            ViewBag.mIntensityGain = 0;
            ViewBag.hIntensityGain = 0;
            return View();
        }
        [HttpPost]
        public IActionResult Index(IndyProject.Models.Calories model)
        {
            model.calcBMR();
            ViewBag.bmr = model.bmr;
            model.calcORM();
            ViewBag.orm = model.orm;
            ViewBag.lIntensityLoss = model.bmr - 250;
            ViewBag.mIntensityLoss = model.bmr - 500;
            ViewBag.hIntensityLoss = model.bmr - 750;
            ViewBag.lIntensityGain = model.bmr + 250;
            ViewBag.mIntensityGain = model.bmr + 500;
            ViewBag.hIntensityGain = model.bmr + 750;
            return View(model);
        }
    }
}
